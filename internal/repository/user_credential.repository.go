package repository

import (
	"context"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserCredentialRepository struct {
	db *pgxpool.Pool
}

type UserCartRole struct {
	UserId   int    `db:"user_id"`
	CartId   int    `db:"cart_id"`
	RoleName string `db:"role_name"`
}

func NewUserCredentialRepository(db *pgxpool.Pool) (*UserCredentialRepository, error) {
	return &UserCredentialRepository{
		db: db,
	}, nil
}

func (u UserCredentialRepository) GetUserCredentialByUserId(userId int) (models.UserCredential, error) {
	sql := `SELECT id, user_id, email, phone, password, created_at, updated_at FROM user_credentials WHERE user_id = $1`
	rows, err := u.db.Query(context.Background(), sql, userId)
	if err != nil {
		return models.UserCredential{}, err
	}

	userCredential, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UserCredential])
	if err != nil {
		return models.UserCredential{}, err
	}

	return userCredential, nil
}

func (u UserCredentialRepository) UpdateUserCredential(newData models.UserCredential) (models.UserCredential, error) {
	sql := `UPDATE user_credentials SET email = $1, phone = $2, password = $3, updated_at = $4 WHERE user_id = $5 RETURNING id, user_id, email, phone, password, created_at, updated_at`

	rows, err := u.db.Query(context.Background(), sql, newData.Email, newData.Phone, newData.Password, time.Now(), newData.UserId)
	if err != nil {
		return models.UserCredential{}, err
	}
	updatedUserCredentials, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UserCredential])
	if err != nil {
		return models.UserCredential{}, err
	}


	return updatedUserCredentials, nil
}

func (u UserCredentialRepository) GetUserCredentialsByEmail(email string) (models.UserCredential, error) {
	sql := `SELECT id, user_id, email, phone, password, created_at, updated_at FROM user_credentials WHERE email = $1`
	rows, err := u.db.Query(context.Background(), sql, email)
	if err != nil {
		return models.UserCredential{}, err
	}
	userCredentials, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UserCredential])
	if err != nil {
		return models.UserCredential{}, err
	}
	return userCredentials, nil
}

func (u UserCredentialRepository) GetUserCartAndRoleByUserId(userId int) (UserCartRole, error) {
	sql := `SELECT users.id AS user_id, carts.id AS cart_id, roles.role_name FROM users JOIN carts ON carts.user_id = users.id JOIN roles ON roles.id = users.role_id WHERE users.id = $1`
	rows, err := u.db.Query(context.Background(), sql, userId)
	if err != nil {
		return UserCartRole{}, err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[UserCartRole])
	if err != nil {
		return UserCartRole{}, err
	}

	return result, nil
}
