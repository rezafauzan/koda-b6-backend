package repository

import (
	"context"
	"math/big"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

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

func (f ForgotPasswordRepository) CreateForgotPasswordData(email string, code_otp *big.Int) (models.ForgotPassword, error) {
	sql := "INSERT INTO forgot_password (email, code_otp, expired_at) VALUES ($1, $2, $3)"
	rows, err := f.db.Query(context.Background(), sql, email, code_otp, time.Now().Add(5 * time.Minute))
	if err != nil {
		return models.ForgotPassword{}, err
	}
	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, err
	}
	return data, nil
}

func (f ForgotPasswordRepository) GetDataByEmailCode(email string, code_otp big.Int) (models.ForgotPassword, error) {
	sql := "SELECT email, code_otp, expired_at FROM forgot_password WHERE email = $1 AND code_otp = $2"
	rows, err := f.db.Query(context.Background(), sql, email, code_otp)
	if err != nil {
		return models.ForgotPassword{}, err
	}
	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, err
	}
	return data, nil
}

func (f ForgotPasswordRepository) DeleteDataByEmailCode(email string, code_otp big.Int) (models.ForgotPassword, error) {
	sql := "DELETE FROM forgot_password WHERE email = $1 AND code_otp = $2"
	rows, err := f.db.Query(context.Background(), sql, email, code_otp)
	if err != nil {
		return models.ForgotPassword{}, err
	}
	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, err
	}
	return data, nil
}