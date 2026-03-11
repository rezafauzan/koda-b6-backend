package repository

import (
	"context"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository() (*UserRepository, error) {
	db, err := lib.DatabaseConnect()
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		db: db,
	}, nil
}

func (u UserRepository) GetUserByEmail(email string) (dto.User, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return dto.User{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.User])
		if err != nil {
		return dto.User{}, err
	}
	
	return user, nil
}
