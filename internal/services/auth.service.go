package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecretKey = "2c9341ca4cf3d87b9e4eb905d6a3ec45" // Test1234 MD5

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

	userId, storedPassword, err := a.userRepo.GetUserCredentialsByEmail(req.Email)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	if req.Password != storedPassword {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : Invalid email or password !")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : " + err.Error())
	}

	return dto.LoginResponseDTO{
		Token:   tokenString,
	}, nil
}
