package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
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

func (u UserCredentialService) GetUserCredentialByUserId(userId int) (dto.UserCredentialResponseWithoutPasswordDTO, error) {
	userCredentials, err := u.userCredentialRepo.GetUserCredentialByUserId(userId)
	if err != nil {
		return dto.UserCredentialResponseWithoutPasswordDTO{}, errors.New("User credential not found !")
	}
	response := dto.UserCredentialResponseWithoutPasswordDTO{
		Id:        userCredentials.Id,
		UserId:    userCredentials.Id,
		Email:     userCredentials.Email,
		Phone:     userCredentials.Phone,
		CreatedAt: userCredentials.CreatedAt,
		UpdatedAt: userCredentials.UpdatedAt,
	}
	return response, nil
}

func (u UserCredentialService) UpdateUserCredential(userId int, newData dto.UpdateUserCredentialDTO) (dto.UserCredentialResponseDTO, error) {
	userCredentials, err := u.userCredentialRepo.GetUserCredentialByUserId(userId)
	if err != nil {
		return dto.UserCredentialResponseDTO{}, err
	}

	if newData.Email == "" {
		newData.Email = userCredentials.Email
	}
	if newData.Phone == "" {
		newData.Phone = userCredentials.Phone
	}
	if newData.Password == "" {
		newData.Password = userCredentials.Password
	} else {
		if len(newData.Password) < 8 {
			return dto.UserCredentialResponseDTO{}, errors.New("Password minimum 8 characters")
		}
		if newData.Password != "" && newData.ConfirmPassword != newData.Password {
			return dto.UserCredentialResponseDTO{}, errors.New("Confirm password not matched")
		}
		hashedPassword, err := lib.HashPassword(newData.Password)
		if err != nil {
			return dto.UserCredentialResponseDTO{}, errors.New("failed to hash password")
		}
		newData.Password = hashedPassword
	}

	if !strings.Contains(newData.Email, "@") {
		return dto.UserCredentialResponseDTO{}, errors.New("Invalid email format")
	}
	if len(newData.Phone) < 10 {
		return dto.UserCredentialResponseDTO{}, errors.New("Phone minimum 10 digits")
	}

	modeledData := models.UserCredential{
		Id:       userCredentials.Id,
		UserId:   userId,
		Email:    newData.Email,
		Phone:    newData.Phone,
		Password: newData.Password,
	}

	updated, err := u.userCredentialRepo.UpdateUserCredential(modeledData)
	if err != nil {
		return dto.UserCredentialResponseDTO{}, err
	}

	response := dto.UserCredentialResponseDTO{
		Id:     updated.Id,
		UserId: updated.UserId,
		Email:  updated.Email,
		Phone:  updated.Phone,
	}
	return response, nil
}
