package repository

import (
	"rezafauzan/koda-b6-golang/internal/lib"

	"github.com/jackc/pgx/v5"
)

type ForgotPasswordRepository struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository() (*ForgotPasswordRepository, error){
	db, err := lib.DatabaseConnect()
	if err != nil {
		return nil, err
	}

	return &ForgotPasswordRepository{
		db: db,
	}, nil
}