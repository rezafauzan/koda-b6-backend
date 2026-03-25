package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
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

func (u RoleService) CreateNewRole(newRole dto.CreateRoleDTO) (dto.RoleResponseDTO, error) {
	if len(newRole.Role_name) < 4 {
		return dto.RoleResponseDTO{}, errors.New("Failed to create role! : Role name length minimum is 4 characters !")
	}
	m := &models.Role{RoleName: newRole.Role_name}
	created, err := u.roleRepo.AddNewRole(m)
	if err != nil {
		return dto.RoleResponseDTO{}, errors.New("Failed to create role: " + err.Error())
	}
	response := dto.RoleResponseDTO{Id: created.Id, Role_name: created.RoleName, Created_at: created.CreatedAt, Updated_at: created.UpdatedAt}

	return response, nil
}

func (u RoleService) GetAllRole() ([]dto.RoleResponseDTO, error) {
	roles, err := u.roleRepo.GetAllRoles()
	if err != nil {
		return []dto.RoleResponseDTO{}, err
	}
	out := make([]dto.RoleResponseDTO, 0, len(roles))
	for _, r := range roles {
		response := dto.RoleResponseDTO{Id: r.Id, Role_name: r.RoleName, Created_at: r.CreatedAt, Updated_at: r.UpdatedAt}

		out = append(out, response)
	}
	return out, nil
}

func (u RoleService) GetRoleById(id int) (dto.RoleResponseDTO, error) {
	role, err := u.roleRepo.GetRoleById(id)
	if err != nil {
		return dto.RoleResponseDTO{}, errors.New("Role not found !")
	}
	response := dto.RoleResponseDTO{Id: role.Id, Role_name: role.RoleName, Created_at: role.CreatedAt, Updated_at: role.UpdatedAt}

	return response, nil
}

func (u RoleService) UpdateRole(newRole dto.UpdateRoleDTO) (dto.RoleResponseDTO, error) {
	if newRole.Role_name != "" && len(newRole.Role_name) < 4 {
		return dto.RoleResponseDTO{}, errors.New("Role name minimum 4 characters")
	}
	m := models.Role{Id: newRole.Id, RoleName: newRole.Role_name}
	updated, err := u.roleRepo.UpdateRole(m)
	if err != nil {
		return dto.RoleResponseDTO{}, err
	}

	response := dto.RoleResponseDTO{Id: updated.Id, Role_name: updated.RoleName, Created_at: updated.CreatedAt, Updated_at: updated.UpdatedAt}

	return response, nil
}

func (u RoleService) DeleteRole(id int) (dto.RoleResponseDTO, error) {
	role, err := u.roleRepo.GetRoleById(id)
	if err != nil {
		return dto.RoleResponseDTO{}, errors.New("Role not found !")
	}

	err = u.roleRepo.DeleteRole(id)

	if err != nil {
		return dto.RoleResponseDTO{}, errors.New("Failed to delete role: " + err.Error())
	}

	response := dto.RoleResponseDTO{Id: role.Id, Role_name: role.RoleName, Created_at: role.CreatedAt, Updated_at: role.UpdatedAt}

	return response, nil
}
