package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserCredentialRepository struct {
	db *pgxpool.Pool
}

func NewUserCredentialRepository(db *pgxpool.Pool) (*UserCredentialRepository, error) {
	return &UserCredentialRepository{
		db: db,
	}, nil
}

func (u UserCredentialRepository) GetAllUserCredentials() ([]models.UserCredential, error) {
	sql := `SELECT id, user_id, email, phone, password, created_at, updated_at FROM user_credentials`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return []models.UserCredential{}, err
	}

	userCredentials, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.UserCredential])
	if err != nil {
		return []models.UserCredential{}, errors.New("Failed to create response get all user credentials! : " + err.Error())
	}

	return userCredentials, nil
}

func (u UserCredentialRepository) GetUserCredentialById(id int) (models.UserCredential, error) {
	sql := `SELECT id, user_id, email, phone, password, created_at, updated_at FROM user_credentials WHERE id = $1`
	rows, err := u.db.Query(context.Background(), sql, id)
	if err != nil {
		return models.UserCredential{}, err
	}

	userCredential, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.UserCredential])
	if err != nil {
		return models.UserCredential{}, err
	}

	return userCredential, nil
}

func (u UserCredentialRepository) UpdateUserCredential(newData models.UserCredential) (models.UserCredential, error) {
	userCredential, err := u.GetUserCredentialById(newData.Id)
	if err != nil {
		return models.UserCredential{}, err
	}

	if newData.Email == "" {
		newData.Email = userCredential.Email
	}

	if newData.Phone == "" {
		newData.Phone = userCredential.Phone
	}

	if newData.Password == "" {
		newData.Password = userCredential.Password
	}

	sql := `UPDATE user_credentials SET user_id = $1, email = $2, phone = $3, password = $4, updated_at = $5 WHERE id = $6`

	_, err = u.db.Exec(context.Background(), sql, newData.UserId, newData.Email, newData.Phone, newData.Password, time.Now(), newData.Id)
	if err != nil {
		return models.UserCredential{}, err
	}

	updatedUserCredential, err := u.GetUserCredentialById(newData.Id)
	if err != nil {
		return models.UserCredential{}, err
	}

	return updatedUserCredential, nil
}
