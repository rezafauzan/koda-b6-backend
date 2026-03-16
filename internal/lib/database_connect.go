package lib

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func DatabaseConnect() (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return nil, errors.New("Failed parse database config!")
	}

	connection, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		return nil, errors.New("Could not connect to database!")
	}
	
	return connection, errors.New("Connection to database established!")
}
