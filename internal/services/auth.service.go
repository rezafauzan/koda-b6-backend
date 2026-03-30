package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"
)

type AuthService struct {
	userCredentialsRepo *repository.UserCredentialRepository
}

func NewAuthService(userCredentialsRepo *repository.UserCredentialRepository) *AuthService {
	return &AuthService{
		userCredentialsRepo: userCredentialsRepo,
	}
}

func (a AuthService) Login(req dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	if !strings.Contains(req.Email, "@") {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : Invalid email format !")
	}

	userCredentials, err := a.userCredentialsRepo.GetUserCredentialsByEmail(req.Email)
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to get user credentials by email : " + err.Error())
	}

	if req.Password != userCredentials.Password {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : Invalid email or password !")
	}

	user, err := a.userCredentialsRepo.GetUserCredentialsByEmail(req.Email)
	
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to get user by email : " + err.Error())
	}

	token, err := lib.GenerateToken(user.UserId)

	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to generate token : " + err.Error())
	}

	return dto.LoginResponseDTO{
		Token: token,
	}, nil
}
