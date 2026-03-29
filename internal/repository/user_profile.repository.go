package repository

import (
	"context"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserProfileRepository struct {
	db *pgxpool.Pool
}

func NewUserProfileRepository(db *pgxpool.Pool) (*UserProfileRepository, error) {
	return &UserProfileRepository{
		db: db,
	}, nil
}

func (u *UserProfileRepository) GetUserProfileByUserId(user_id int) (models.UserProfile, error) {
	sql := `SELECT id, user_id, user_avatar, first_name, last_name, address, created_at, updated_at FROM user_profiles WHERE user_id = $1`
	rows, err := u.db.Query(context.Background(), sql, user_id)
	if err != nil {
		return models.UserProfile{}, err
	}

	userProfile, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UserProfile])
	if err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (u UserProfileRepository) UpdateUserProfile(newData models.UserProfile) (models.UserProfile, error) {
	sql := `UPDATE user_profiles SET user_avatar = $1, first_name = $2, last_name = $3, address = $4, updated_at = $5 WHERE user_id = $6`

	_, err := u.db.Exec(context.Background(), sql, newData.UserAvatar, newData.FirstName, newData.LastName, newData.Address, time.Now(), newData.UserId)
	if err != nil {
		return models.UserProfile{}, err
	}

	updatedUserProfile, err := u.GetUserProfileByUserId(newData.UserId)
	if err != nil {
		return models.UserProfile{}, err
	}

	return updatedUserProfile, nil
}