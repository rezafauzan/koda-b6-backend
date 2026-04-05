package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (u UserRepository) CreateNewUser(newUser dto.CreateUserDTO) (dto.CreateUserDTO, error) {
	sql := "INSERT INTO users (role_id,verified,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id, role_id, verified, created_at,updated_at"
	rows, err := u.db.Query(context.Background(), sql, 2, true, time.Now(), time.Now())
	if err != nil {
		return dto.CreateUserDTO{}, errors.New("Failed to create new user! : " + err.Error())
	}

	registeredUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return dto.CreateUserDTO{}, errors.New("Failed to create new user! : " + err.Error())
	}

	sql = "INSERT INTO user_credentials (user_id, email, phone, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = u.db.Exec(context.Background(), sql, registeredUser.Id, newUser.Email, newUser.Phone, newUser.Password, time.Now(), time.Now())

	if err != nil {
		return dto.CreateUserDTO{}, errors.New("Failed to create new user! : " + err.Error())
	}

	sql = "INSERT INTO user_profiles (user_id, user_avatar, first_name, last_name, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = u.db.Exec(context.Background(), sql, registeredUser.Id, "https://i.pravatar.cc/400?img=4", newUser.First_name, newUser.Last_name, newUser.Address, time.Now(), time.Now())

	if err != nil {
		return dto.CreateUserDTO{}, errors.New("Failed to create new user! : " + err.Error())
	}

	return newUser, nil
}

func (u UserRepository) GetAllUsers() ([]dto.UserResponseDTO, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return []dto.UserResponseDTO{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.UserResponseDTO])
	if err != nil {
		return []dto.UserResponseDTO{}, errors.New("Failed to create response get all users! : " + err.Error())
	}

	return users, nil
}

func (u UserRepository) GetUserByEmail(email string) (dto.UserResponseDTO, bool, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id WHERE user_credentials.email = $1`
	rows, err := u.db.Query(context.Background(), sql, email)
	if err != nil {
		return dto.UserResponseDTO{}, false, errors.New("Failed to fetch user by email from database : " + err.Error())
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.UserResponseDTO])
	if err != nil {
		return dto.UserResponseDTO{}, false, errors.New("Failed to convert rows to struct : " + err.Error())
	}
	return user, true, nil
}

func (u UserRepository) GetUserByPhone(phone string) (dto.UserResponseDTO, bool, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id WHERE user_credentials.phone = $1`
	rows, err := u.db.Query(context.Background(), sql, phone)
	if err != nil {
		return dto.UserResponseDTO{}, false, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.UserResponseDTO])
	if err != nil {
		return dto.UserResponseDTO{}, false, err
	}

	return user, true, nil
}

func (u UserRepository) GetUserById(id int) (dto.UserResponseDTO, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id WHERE users.id = $1`
	rows, err := u.db.Query(context.Background(), sql, id)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.UserResponseDTO])
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	return user, nil
}

func (u UserRepository) UpdateUserProfile(newData dto.UpdateUserProfileDTO) (dto.UserResponseDTO, error) {
	user, err := u.GetUserById(newData.Id)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	if newData.First_name == "" {
		newData.First_name = user.First_name
	}

	if newData.Last_name == "" {
		newData.Last_name = user.Last_name
	}

	if newData.Address == "" {
		newData.Address = user.Address
	}

	if newData.User_avatar == "" {
		newData.User_avatar = user.User_avatar
	}

	sql := `UPDATE user_profiles SET first_name = $1, last_name = $2, address = $3, user_avatar = $4, updated_at = $5 WHERE user_id = $6`

	_, err = u.db.Exec(context.Background(), sql, newData.First_name, newData.Last_name, newData.Address, newData.User_avatar, time.Now(), newData.Id)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	updatedUser, err := u.GetUserById(newData.Id)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	return updatedUser, nil
}

func (u UserRepository) DeleteUser(id int) error {
	sql := `DELETE FROM user_credentials WHERE user_id = $1`
	_, err := u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	sql = `DELETE FROM user_profiles WHERE user_id = $1`
	_, err = u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	sql = `DELETE FROM users WHERE id = $1`
	_, err = u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) UpdatePassword(email string, newPassword string) error {
	sql := `UPDATE user_credentials SET password = $1, updated_at = $2 WHERE email = $3`
	_, err := u.db.Exec(context.Background(), sql, newPassword, time.Now(), email)
	if err != nil {
		return err
	}

	return nil
}
