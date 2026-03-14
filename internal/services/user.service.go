package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func (u UserService) NewUserService() *UserService {
	return &UserService{
		userRepo: &repository.UserRepository{},
	}
}

func (u UserService) AddNewUser(newUser dto.UserRegister) (dto.UserRegister, error){
	if len(newUser.First_name) < 4 {
	}
	if len(newUser.Last_name) < 4 {
		return dto.UserRegister{}, errors.New("Failed to create user! : Last name length minimum is 4 characters !")
	}
	if !strings.Contains(newUser.Email, "@") {
		return dto.UserRegister{}, errors.New("Failed to create user! : Invalid email format !")
	}
	if len(newUser.Phone) < 10 {
		return dto.UserRegister{}, errors.New("Failed to create user! : Phone numbers length minimum 10 digits !")
	}
	if len(newUser.Address) < 10 {
		return dto.UserRegister{}, errors.New("Failed to create user! : Address length minimum is 10 characters !")
	}
	if len(newUser.Password) < 8 {
		return dto.UserRegister{}, errors.New("Failed to create user! : Password too weak minimum length is 8 characters !")
	}
	if newUser.Password_confirm != newUser.Password {
		return dto.UserRegister{}, errors.New("Failed to create user! : Password confirmation missmatch !")
	}
	return u.userRepo.AddNewUser(newUser)
}
