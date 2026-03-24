package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ForgotPasswordRepository struct {
	db *pgxpool.Pool
}

func NewForgotPasswordRepository(db *pgxpool.Pool) (*ForgotPasswordRepository, error) {
	return &ForgotPasswordRepository{
		db: db,
	}, nil
}

func (u ForgotPasswordRepository) CreateForgotPassword(newForgotPassword *models.ForgotPassword) (models.ForgotPassword, error) {
	sql := "INSERT INTO forgot_password (email, code_otp, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, email, code_otp, created_at, updated_at"
	rows, err := u.db.Query(context.Background(), sql, newForgotPassword.Email, newForgotPassword.CodeOtp, time.Now(), time.Now())
	if err != nil {
		return models.ForgotPassword{}, errors.New("Failed to create new forgot password! : " + err.Error())
	}

	registeredForgotPassword, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, errors.New("Failed to create new forgot password! : " + err.Error())
	}

	return registeredForgotPassword, nil
}

func (u ForgotPasswordRepository) GetAllForgotPasswords() ([]models.ForgotPassword, error) {
	sql := `SELECT id, email, code_otp, created_at, updated_at FROM forgot_password`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return []models.ForgotPassword{}, err
	}

	forgotPasswords, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return []models.ForgotPassword{}, errors.New("Failed to create response get all forgot passwords! : " + err.Error())
	}

	return forgotPasswords, nil
}

func (u ForgotPasswordRepository) GetForgotPasswordById(id int) (models.ForgotPassword, error) {
	sql := `SELECT id, email, code_otp, created_at, updated_at FROM forgot_password WHERE id = $1`
	rows, err := u.db.Query(context.Background(), sql, id)
	if err != nil {
		return models.ForgotPassword{}, err
	}

	forgotPassword, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, err
	}

	return forgotPassword, nil
}

func (u ForgotPasswordRepository) UpdateForgotPassword(newData models.ForgotPassword) (models.ForgotPassword, error) {
	forgotPassword, err := u.GetForgotPasswordById(newData.Id)
	if err != nil {
		return models.ForgotPassword{}, err
	}

	if newData.Email == "" {
		newData.Email = forgotPassword.Email
	}

	sql := `UPDATE forgot_password SET email = $1, code_otp = $2, updated_at = $3 WHERE id = $4`

	_, err = u.db.Exec(context.Background(), sql, newData.Email, newData.CodeOtp, time.Now(), newData.Id)
	if err != nil {
		return models.ForgotPassword{}, err
	}

	updatedForgotPassword, err := u.GetForgotPasswordById(newData.Id)
	if err != nil {
		return models.ForgotPassword{}, err
	}

	return updatedForgotPassword, nil
}

func (u ForgotPasswordRepository) DeleteForgotPassword(id int) error {
	sql := `DELETE FROM forgot_password WHERE id = $1`
	_, err := u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (u ForgotPasswordRepository) ClearForgotPassword(email string) error {
	sql := `DELETE FROM forgot_password WHERE email = $1`
	_, err := u.db.Exec(context.Background(), sql, email)
	if err != nil {
		return err
	}

	return nil
}

func (u ForgotPasswordRepository) CreateForgotPasswordData(email string, code_otp int) (models.ForgotPassword, error) {
	sql := "INSERT INTO forgot_password (email, code_otp, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, email, code_otp, created_at, updated_at"
	rows, err := u.db.Query(context.Background(), sql, email, code_otp, time.Now(), time.Now())
	if err != nil {
		return models.ForgotPassword{}, errors.New("Failed to create new forgot password! : " + err.Error())
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, errors.New("Failed to create new forgot password! : " + err.Error())
	}

	return data, nil
}

func (u ForgotPasswordRepository) GetLatestOTP(email string) (models.ForgotPassword, error) {
	sql := "SELECT id, email, code_otp, created_at, updated_at FROM forgot_password WHERE email = $1 ORDER BY created_at DESC, id DESC LIMIT 1"
	rows, err := u.db.Query(context.Background(), sql, email)
	if err != nil {
		return models.ForgotPassword{}, err
	}

	data, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return models.ForgotPassword{}, err
	}

	return data, nil
}

func (u ForgotPasswordRepository) MarkOTPUsed(id int) error {
	sql := "DELETE FROM forgot_password WHERE id = $1"
	_, err := u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}
