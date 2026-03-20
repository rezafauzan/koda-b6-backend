package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (a AuthService) Login(req dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	if !strings.Contains(req.Email, "@") {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : Invalid email format !")
	}

	userCred, err := a.userRepo.GetUserCredentialsByEmail(req.Email)
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to get user credentials by email : " + err.Error())
	}

	if req.Password != userCred.Password {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : Invalid email or password !")
	}

	user, _, err := a.userRepo.GetUserByEmail(req.Email)
	
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to get user by email : " + err.Error())
	}

	token, err := lib.GenerateToken(user)

	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to generate token : " + err.Error())
	}

	return dto.LoginResponseDTO{
		Token: token,
	}, nil
}
