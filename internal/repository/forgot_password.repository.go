package repository

import (
	"context"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/models"

	"github.com/jackc/pgx/v5"
)

type ForgotPasswordRepository struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository() (*ForgotPasswordRepository, error) {
	db, err := lib.DatabaseConnect()
	if err != nil {
		return nil, err
	}

	return &ForgotPasswordRepository{
		db: db,
	}, nil
}

func (f ForgotPasswordRepository) GetDataByEmailCode(email string) (models.ForgotPassword, error) {
	sql := "SELECT email, code_otp, expired_at, created_at, updated_at FROM forgot_password WHERE email = $1"
	rows, err := f.db.Query(context.Background(), sql, email)
	if err != nil {
		return models.ForgotPassword{}, err
	}
	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, err
	}
	return data, nil
}
