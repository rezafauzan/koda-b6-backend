package repository

import (
	"context"
	"errors"
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

func (u UserProfileRepository) CreateUserProfile(newUserProfile *models.UserProfile) (models.UserProfile, error) {
	sql := "INSERT INTO user_profiles (user_id, user_avatar, first_name, last_name, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, user_id, user_avatar, first_name, last_name, address, created_at, updated_at"
	rows, err := u.db.Query(context.Background(), sql, newUserProfile.UserId, newUserProfile.UserAvatar, newUserProfile.FirstName, newUserProfile.LastName, newUserProfile.Address, time.Now(), time.Now())
	if err != nil {
		return models.UserProfile{}, errors.New("Failed to create new user profile! : " + err.Error())
	}

	registeredUserProfile, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UserProfile])
	if err != nil {
		return models.UserProfile{}, errors.New("Failed to create new user profile! : " + err.Error())
	}

	return registeredUserProfile, nil
}

func (u UserProfileRepository) GetAllUserProfiles() ([]models.UserProfile, error) {
	sql := `SELECT id, user_id, user_avatar, first_name, last_name, address, created_at, updated_at FROM user_profiles`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return []models.UserProfile{}, err
	}

	userProfiles, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.UserProfile])
	if err != nil {
		return []models.UserProfile{}, errors.New("Failed to create response get all user profiles! : " + err.Error())
	}

	return userProfiles, nil
}

func (u UserProfileRepository) GetUserProfileById(id int) (models.UserProfile, error) {
	sql := `SELECT id, user_id, user_avatar, first_name, last_name, address, created_at, updated_at FROM user_profiles WHERE id = $1`
	rows, err := u.db.Query(context.Background(), sql, id)
	if err != nil {
		return models.UserProfile{}, err
	}

	userProfile, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.UserProfile])
	if err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (u UserProfileRepository) UpdateUserProfile(newData models.UserProfile) (models.UserProfile, error) {
	userProfile, err := u.GetUserProfileById(newData.Id)
	if err != nil {
		return models.UserProfile{}, err
	}

	if newData.UserAvatar == "" {
		newData.UserAvatar = userProfile.UserAvatar
	}

	if newData.FirstName == "" {
		newData.FirstName = userProfile.FirstName
	}

	if newData.LastName == "" {
		newData.LastName = userProfile.LastName
	}

	if newData.Address == "" {
		newData.Address = userProfile.Address
	}

	sql := `UPDATE user_profiles SET user_id = $1, user_avatar = $2, first_name = $3, last_name = $4, address = $5, updated_at = $6 WHERE id = $7`

	_, err = u.db.Exec(context.Background(), sql, newData.UserId, newData.UserAvatar, newData.FirstName, newData.LastName, newData.Address, time.Now(), newData.Id)
	if err != nil {
		return models.UserProfile{}, err
	}

	updatedUserProfile, err := u.GetUserProfileById(newData.Id)
	if err != nil {
		return models.UserProfile{}, err
	}

	return updatedUserProfile, nil
}

func (u UserProfileRepository) DeleteUserProfile(id int) error {
	sql := `DELETE FROM user_profiles WHERE id = $1`
	_, err := u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}
