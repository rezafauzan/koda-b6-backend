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

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u UserService) AddNewUser(newUser *dto.CreateUserDTO) (*dto.CreateUserDTO, error) {
	if len(newUser.First_name) < 4 {
		return &dto.CreateUserDTO{}, errors.New("Failed to create user! : First name length minimum is 4 characters !")
	}
	if len(newUser.Last_name) < 4 {
		return &dto.CreateUserDTO{}, errors.New("Failed to create user! : Last name length minimum is 4 characters !")
	}
	if !strings.Contains(newUser.Email, "@") {
		return &dto.CreateUserDTO{}, errors.New("Failed to create user! : Invalid email format !")
	}
	if len(newUser.Phone) < 10 {
		return &dto.CreateUserDTO{}, errors.New("Failed to create user! : Phone numbers length minimum 10 digits !")
	}
	if len(newUser.Address) < 10 {
		return &dto.CreateUserDTO{}, errors.New("Failed to create user! : Address length minimum is 10 characters !")
	}
	if len(newUser.Password) < 8 {
		return &dto.CreateUserDTO{}, errors.New("Failed to create user! : Password too weak minimum length is 8 characters !")
	}
	if newUser.Password_confirm != newUser.Password {
		return &dto.CreateUserDTO{}, errors.New("Failed to create user! : Password confirmation missmatch !")
	}
	return u.userRepo.AddNewUser(newUser)
}

func (u UserService) GetAllUser() ([]dto.UserResponseDTO, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return []dto.UserResponseDTO{}, err
	}

	return users, nil
}

func (u UserService) UpdateUserProfiles(newUser dto.UpdateUserProfileDTO) (dto.UserResponseDTO, error) {

	if newUser.First_name != "" && len(newUser.First_name) < 4 {
		return dto.UserResponseDTO{}, errors.New("First name minimum 4 characters")
	}

	if newUser.Last_name != "" && len(newUser.Last_name) < 4 {
		return dto.UserResponseDTO{}, errors.New("Last name minimum 4 characters")
	}

	if newUser.Address != "" && len(newUser.Address) < 10 {
		return dto.UserResponseDTO{}, errors.New("Address minimum 10 characters")
	}

	if newUser.User_avatar != "" && len(newUser.User_avatar) < 10 {
		return dto.UserResponseDTO{}, errors.New("User avatar minimum 10 characters")
	}

	return u.userRepo.UpdateUserProfile(newUser)
}

func (u UserService) DeleteUser(id int) (dto.UserResponseDTO, error) {
	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		return user, errors.New("User not found !")
	}

	err = u.userRepo.DeleteUser(id)

	if err != nil {
		return dto.UserResponseDTO{}, errors.New("Failed to delete user: " + err.Error())
	}

	return user, nil
}
