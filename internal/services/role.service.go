package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/models"
	"rezafauzan/koda-b6-golang/internal/repository"
)

type RoleService struct {
	roleRepo *repository.RoleRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

func (u RoleService) AddNewRole(newRole models.Role) (models.Role, error) {
	if len(newRole.Role_name) < 4 {
		return models.Role{}, errors.New("Failed to create role! : Role name length minimum is 4 characters !")
	}
	return u.roleRepo.AddNewRole(newRole)
}

func (u RoleService) GetAllRole() ([]models.Role, error) {
	roles, err := u.roleRepo.GetAllRoles()
	if err != nil {
		return []models.Role{}, err
	}

	return roles, nil
}

func (u RoleService) UpdateRoles(newRole models.Role) (models.Role, error) {

	if newRole.Role_name != "" && len(newRole.Role_name) < 4 {
		return models.Role{}, errors.New("Role name minimum 4 characters")
	}

	return u.roleRepo.UpdateRole(newRole)
}

func (u RoleService) DeleteRole(id int) (models.Role, error) {
	role, err := u.roleRepo.GetRoleById(id)
	if err != nil {
		return role, errors.New("Role not found !")
	}

	err = u.roleRepo.DeleteRole(id)

	if err != nil {
		return models.Role{}, errors.New("Failed to delete role: " + err.Error())
	}

	return role, nil
}
