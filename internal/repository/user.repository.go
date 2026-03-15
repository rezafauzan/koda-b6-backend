package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (u UserRepository) AddNewUser(newUser *dto.UserRegister) (*dto.UserRegister, error){
	sql := "INSERT INTO users (role_id,verified,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id, role_id, verified, created_at,updated_at"
	rows, err := u.db.Query(context.Background(), sql, 2, false, time.Now(), time.Now())
	if err != nil {
		return &dto.UserRegister{}, errors.New("Failed to create new user! : " + err.Error())
	}

	registeredUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return &dto.UserRegister{}, errors.New("Failed to create new user! : " + err.Error())
	}

	sql = "INSERT INTO user_profiles (user_id, user_avatar, first_name, last_name, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = u.db.Exec(context.Background(), sql, registeredUser.Id, "https://i.pravatar.cc/400?img=4", newUser.First_name, newUser.Last_name, newUser.Address, time.Now(), time.Now())

	if err != nil {
		return &dto.UserRegister{}, errors.New("Failed to create new user! : " + err.Error())
	}

	sql = "INSERT INTO user_credentials (user_id, email, phone, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = u.db.Exec(context.Background(), sql, registeredUser.Id, newUser.Email, newUser.Phone, newUser.Password, time.Now(), time.Now())

	if err != nil {
		return &dto.UserRegister{}, errors.New("Failed to create new user! : " + err.Error())
	}
	return newUser, nil
}

func (u UserRepository) GetAllUsers() ([]dto.User, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return []dto.User{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.User])
	if err != nil {
		return []dto.User{}, errors.New("Failed to create response get all users! : " + err.Error())
	}

	return users, nil
}

func (u UserRepository) GetUserByEmail(email string) (dto.User, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id WHERE user_credentials.email = '$1'`
	rows, err := u.db.Query(context.Background(), sql, email)
	if err != nil {
		return dto.User{}, err
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[dto.User])
		if err != nil {
		return dto.User{}, err
	}
	
	return user, nil
}

func (u UserRepository) GetUserById(id int) (dto.User, error) {
	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id WHERE users.id = '$1'`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return dto.User{}, err
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[dto.User])
		if err != nil {
		return dto.User{}, err
	}
	
	return user, nil
}