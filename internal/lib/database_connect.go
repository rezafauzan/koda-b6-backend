package lib

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DatabaseConnect() (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")

	var pool *pgxpool.Pool
	var err error

	for i := 0; i < 10; i++ {
		pool, err = pgxpool.New(context.Background(), dbURL)
		if err == nil {
			// cek koneksi beneran
			err = pool.Ping(context.Background())
			if err == nil {
				fmt.Println("Connection to database established!")
				return pool, nil
			}
		}

		fmt.Println("Database not ready, retrying...")
		time.Sleep(4 * time.Second)
	}

	return nil, errors.New("could not connect to database: " + err.Error())
}
