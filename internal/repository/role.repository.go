package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type RoleRepository struct {
	db *pgx.Conn
}

func NewRoleRepository(db *pgx.Conn) (*RoleRepository, error) {
	return &RoleRepository{
		db: db,
	}, nil
}

func (u RoleRepository) AddNewRole(newRole models.Role) (models.Role, error) {
	sql := "INSERT INTO roles (role_name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id, role_name, created_at, updated_at"
	rows, err := u.db.Query(context.Background(), sql, newRole.Role_name, time.Now(), time.Now())
	if err != nil {
		return models.Role{}, errors.New("Failed to create new role! : " + err.Error())
	}

	registeredRole, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Role])
	if err != nil {
		return models.Role{}, errors.New("Failed to create new role! : " + err.Error())
	}

	return registeredRole, nil
}

func (u RoleRepository) GetAllRoles() ([]models.Role, error) {
	sql := `SELECT id, role_name, created_at, updated_at FROM roles`
	rows, err := u.db.Query(context.Background(), sql)
	if err != nil {
		return []models.Role{}, err
	}

	roles, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Role])
	if err != nil {
		return []models.Role{}, errors.New("Failed to create response get all roles! : " + err.Error())
	}

	return roles, nil
}

func (u RoleRepository) GetRoleById(id int) (models.Role, error) {
	sql := `SELECT id, role_name, created_at, updated_at FROM roles WHERE id = $1`
	rows, err := u.db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Role{}, err
	}

	role, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.Role])
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}