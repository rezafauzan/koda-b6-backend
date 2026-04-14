package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"

	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	userCredentialsRepo *repository.UserCredentialRepository
	userRepo            *repository.UserRepository
	cartItemRepo        *repository.CartItemRepository
}

func NewAuthService(userCredentialsRepo *repository.UserCredentialRepository, userRepo *repository.UserRepository, cartItemRepo *repository.CartItemRepository) *AuthService {
	return &AuthService{
		userCredentialsRepo: userCredentialsRepo,
		userRepo:            userRepo,
		cartItemRepo:        cartItemRepo,
	}
}

func (a AuthService) Login(req dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	if !strings.Contains(req.Email, "@") {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : Invalid email format !")
	}

	userCredentials, err := a.userCredentialsRepo.GetUserCredentialsByEmail(req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.LoginResponseDTO{}, errors.New("Invalid email or password !")
		}
		return dto.LoginResponseDTO{}, errors.New("failed to get user credentials by email: " + err.Error())
	}

	if req.Password != userCredentials.Password {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! : Invalid email or password !")
	}

	user, err := a.userCredentialsRepo.GetUserCredentialsByEmail(req.Email)
	userCart, err := a.cartItemRepo.GetCartByUserId(user.Id)

	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to get user by email : " + err.Error())
	}

	token, err := lib.GenerateToken(user.UserId, userCart.UserId)

	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to generate token : " + err.Error())
	}

	return dto.LoginResponseDTO{
		Token: token,
	}, nil
}

func (a AuthService) Register(newUser dto.CreateUserDTO) (dto.CreateUserDTO, error) {
	if len(newUser.First_name) < 4 {
		return dto.CreateUserDTO{}, errors.New("Failed to create user! : First name length minimum is 4 characters !")
	}
	if len(newUser.Last_name) < 4 {
		return dto.CreateUserDTO{}, errors.New("Failed to create user! : Last name length minimum is 4 characters !")
	}
	if !strings.Contains(newUser.Email, "@") {
		return dto.CreateUserDTO{}, errors.New("Failed to create user! : Invalid email format !")
	}
	if len(newUser.Phone) < 10 {
		return dto.CreateUserDTO{}, errors.New("Failed to create user! : Phone numbers length minimum 10 digits !")
	}
	if len(newUser.Address) < 10 {
		return dto.CreateUserDTO{}, errors.New("Failed to create user! : Address length minimum is 10 characters !")
	}
	if len(newUser.Password) < 8 {
		return dto.CreateUserDTO{}, errors.New("Failed to create user! : Password too weak minimum length is 8 characters !")
	}
	if newUser.Password_confirm != newUser.Password {
		return dto.CreateUserDTO{}, errors.New("Failed to create user! : Password confirmation missmatch !")
	}

	_, emailExist, _ := a.userRepo.GetUserByEmail(newUser.Email)

	if emailExist {
		return dto.CreateUserDTO{}, errors.New("Failed to create new user! : Email allready used !")
	}

	_, phoneExist, _ := a.userRepo.GetUserByPhone(newUser.Phone)

	if phoneExist {
		return dto.CreateUserDTO{}, errors.New("Failed to create new user! : Phone number allready used !")
	}

	return a.userRepo.CreateNewUser(newUser)
}
