package lib

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DatabaseConnect() (*pgx.Conn, error) {
	connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, errors.New("Could not connect to database!")
	}
	fmt.Println("Connection to database established!")
	return connection, nil
}
