package openapi

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type ToyAPIDB interface {
	GetData(ctx context.Context) ([]DataOutput, error)
	AddData(ctx context.Context, name string, quantity int) (DataOutput, error)
	GetDataByID(ctx context.Context, id int) (DataOutput, error)
	UpdateDataByID(ctx context.Context, id int, name string, quantity int) (DataOutput, error)
	DeleteDataByID(ctx context.Context, id int) error
	Close() error
}

type APIDatabase struct {
	db *sql.DB
}

func NewDBConn(host string, port int, user string, password string, dbName string) (*APIDatabase, error) {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &APIDatabase{db}, nil
}

func (db *APIDatabase) GetData(ctx context.Context) ([]DataOutput, error) {
	result, err := db.db.QueryContext(ctx, "SELECT id, name, quantity FROM data_table")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	returnData := make([]DataOutput, 0)
	for result.Next() {
		var row DataOutput
		if err = result.Scan(&row.DataId, &row.Name, &row.Quantity); err != nil {
			return nil, err
		}
		returnData = append(returnData, row)
	}
	return returnData, nil
}

func (db *APIDatabase) AddData(ctx context.Context, name string, quantity int) (DataOutput, error) {
	const query = "INSERT INTO data_table (name, quantity) VALUES ($1, $2) RETURNING id"
	var id int
	err := db.db.QueryRowContext(ctx, query, name, quantity).Scan(&id)
	if err != nil {
		return DataOutput{}, err
	}
	return db.GetDataByID(ctx, int(id))
}
func (db *APIDatabase) GetDataByID(ctx context.Context, id int) (DataOutput, error) {
	const query = "SELECT id, name, quantity FROM data_table WHERE id = $1"
	result := db.db.QueryRowContext(ctx, query, id)
	var returnVal DataOutput
	if err := result.Scan(&returnVal.DataId, &returnVal.Name, &returnVal.Quantity); err != nil {
		return returnVal, err
	}
	return returnVal, nil

}
func (db *APIDatabase) UpdateDataByID(ctx context.Context, id int, name string, quantity int) (DataOutput, error) {
	updateParams := make([]string, 0, 2)
	updateParamValues := make([]interface{}, 0, 2)
	if name != "" {
		updateParams = append(updateParams, "name")
		updateParamValues = append(updateParamValues, name)
	}
	if quantity != 0 {
		updateParams = append(updateParams, "quantity")
		updateParamValues = append(updateParamValues, quantity)
	}
	var updateStr []string
	for i := range updateParams {
		updateStr = append(updateStr, fmt.Sprintf("%s = $%d", updateParams[i], i+1))
	}
	query := fmt.Sprintf("UPDATE data_table SET %s WHERE id = $%d", strings.Join(updateStr, ","), len(updateParams)+1)
	updateParamValues = append(updateParamValues, id)
	result, err := db.db.ExecContext(ctx, query, updateParamValues...)
	if err != nil {
		return DataOutput{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return DataOutput{}, err
	}
	if rowsAffected != 1 {
		return DataOutput{}, fmt.Errorf("no such id: %d", id)
	}
	return db.GetDataByID(ctx, id)
}
func (db *APIDatabase) DeleteDataByID(ctx context.Context, id int) error {
	const query = "DELETE FROM data_table WHERE id = $1"
	result, err := db.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("no such id: %d", id)
	}
	return nil
}
func (db *APIDatabase) Close() error {
	return db.db.Close()
}
