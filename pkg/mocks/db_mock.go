package mocks

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
)

type DBMock struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

func NewDBMock() (*DBMock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	return &DBMock{
		DB:   db,
		Mock: mock,
	}, nil
}
