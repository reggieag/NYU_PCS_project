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
	UpdateDataByID(ctx context.Context, id int, name *string, quantity *int) (DataOutput, error)
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
		return nil, fmt.Errorf("unable to create db connection: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to open db connection: %w", err)
	}
	return &APIDatabase{db}, nil
}

func (db *APIDatabase) GetData(ctx context.Context) ([]DataOutput, error)
func (db *APIDatabase) AddData(ctx context.Context, name string, quantity int) (DataOutput, error)
func (db *APIDatabase) GetDataByID(ctx context.Context, id int) (DataOutput, error)
func (db *APIDatabase) UpdateDataByID(ctx context.Context, id int, name *string, quantity *int) (DataOutput, error)
func (db *APIDatabase) DeleteDataByID(ctx context.Context, id int) error
func (db *APIDatabase) Close() error {
	return db.db.Close()
}
