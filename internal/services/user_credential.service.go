package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"
)

type UserCredentialService struct {
	userCredentialRepo *repository.UserCredentialRepository
}

func NewUserCredentialService(userCredentialRepo *repository.UserCredentialRepository) *UserCredentialService {
	return &UserCredentialService{
		userCredentialRepo: userCredentialRepo,
	}
}

func (u UserCredentialService) GetAllUserCredential() ([]dto.UserCredentialResponseDTO, error) {
	list, err := u.userCredentialRepo.GetAllUserCredentials()
	if err != nil {
		return []dto.UserCredentialResponseDTO{}, err
	}
	out := make([]dto.UserCredentialResponseDTO, 0, len(list))
	for _, x := range list {
		out = append(out, dto.UserCredentialResponseFromModel(x))
	}
	return out, nil
}

func (u UserCredentialService) GetUserCredentialById(id int) (dto.UserCredentialResponseDTO, error) {
	x, err := u.userCredentialRepo.GetUserCredentialById(id)
	if err != nil {
		return dto.UserCredentialResponseDTO{}, errors.New("User credential not found !")
	}
	return dto.UserCredentialResponseFromModel(x), nil
}

func (u UserCredentialService) UpdateUserCredential(newData dto.UpdateUserCredentialDTO) (dto.UserCredentialResponseDTO, error) {
	if newData.Email != "" && !strings.Contains(newData.Email, "@") {
		return dto.UserCredentialResponseDTO{}, errors.New("Invalid email format")
	}
	if newData.Phone != "" && len(newData.Phone) < 10 {
		return dto.UserCredentialResponseDTO{}, errors.New("Phone minimum 10 digits")
	}
	if newData.Password != "" && len(newData.Password) < 8 {
		return dto.UserCredentialResponseDTO{}, errors.New("Password minimum 8 characters")
	}

	m := models.UserCredential{
		Id: newData.Id, UserId: newData.User_id, Email: newData.Email, Phone: newData.Phone, Password: newData.Password,
	}
	updated, err := u.userCredentialRepo.UpdateUserCredential(m)
	if err != nil {
		return dto.UserCredentialResponseDTO{}, err
	}
	return dto.UserCredentialResponseFromModel(updated), nil
}