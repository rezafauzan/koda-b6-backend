package lib

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DatabaseConnect() (*pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env!")
	}

	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		fmt.Println("Failed parse database config!")
	}

	connection, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		fmt.Println("Could not connect to database!")
	} else {
		fmt.Println("Connection to database established!")
	}

	return connection, err
}
