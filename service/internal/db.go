package openapi

import (
	"context"
	"database/sql"
	"fmt"

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

/*
* All these queries are exactly the wrong way to be doing SQL.
* This is the point...
 */

func (db *APIDatabase) GetData(ctx context.Context) ([]DataOutput, error) {
	query := fmt.Sprintf("SELECT id, name, quantity FROM data_table")
	result, err := db.db.QueryContext(ctx, query)
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
	query := fmt.Sprintf("INSERT INTO data_table (name, quantity) VALUES ('%s', %d) RETURNING id", name, quantity)
	var id int
	err := db.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return DataOutput{}, err
	}
	return db.GetDataByID(ctx, int(id))
}
func (db *APIDatabase) GetDataByID(ctx context.Context, id int) (DataOutput, error) {
	query := fmt.Sprintf("SELECT id, name, quantity FROM data_table WHERE id = %d", id)
	result := db.db.QueryRowContext(ctx, query)
	var returnVal DataOutput
	if err := result.Scan(&returnVal.DataId, &returnVal.Name, &returnVal.Quantity); err != nil {
		return returnVal, err
	}
	return returnVal, nil

}
func (db *APIDatabase) UpdateDataByID(ctx context.Context, id int, name string, quantity int) (DataOutput, error) {
	return DataOutput{}, nil
}
func (db *APIDatabase) DeleteDataByID(ctx context.Context, id int) error {
	return nil
}
func (db *APIDatabase) Close() error {
	return db.db.Close()
}
