package lib

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func DatabaseConnect() (*pgx.Conn, error) {
	dbURL := os.Getenv("DATABASE_URL")

	var connection *pgx.Conn
	var err error

	for i := 0; i < 10; i++ {
		connection, err = pgx.Connect(context.Background(), dbURL)
		if err == nil {
			fmt.Println("Connection to database established!")
			return connection, nil
		}

		fmt.Println("Database not ready, retrying...")
		time.Sleep(4 * time.Second)
	}

	return nil, errors.New("could not connect to database : " + err.Error())
}
