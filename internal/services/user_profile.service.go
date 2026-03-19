package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"
	"rezafauzan/koda-b6-golang/internal/repository"
)

type UserProfileService struct {
	userProfileRepo *repository.UserProfileRepository
}

func NewUserProfileService(userProfileRepo *repository.UserProfileRepository) *UserProfileService {
	return &UserProfileService{
		userProfileRepo: userProfileRepo,
	}
}

func (u UserProfileService) CreateNewUserProfile(newData dto.CreateUserProfileDTO) (dto.UserProfileResponseDTO, error) {
	if newData.User_id <= 0 {
		return dto.UserProfileResponseDTO{}, errors.New("Failed to create user profile! : User id is required !")
	}
	if len(newData.First_name) < 4 {
		return dto.UserProfileResponseDTO{}, errors.New("Failed to create user profile! : First name length minimum is 4 characters !")
	}
	if len(newData.Last_name) < 4 {
		return dto.UserProfileResponseDTO{}, errors.New("Failed to create user profile! : Last name length minimum is 4 characters !")
	}
	if len(newData.Address) < 10 {
		return dto.UserProfileResponseDTO{}, errors.New("Failed to create user profile! : Address length minimum is 10 characters !")
	}
	if newData.User_avatar != "" && len(newData.User_avatar) < 10 {
		return dto.UserProfileResponseDTO{}, errors.New("Failed to create user profile! : User avatar length minimum is 10 characters !")
	}

	m := &models.UserProfile{
		UserId: newData.User_id, UserAvatar: newData.User_avatar, FirstName: newData.First_name,
		LastName: newData.Last_name, Address: newData.Address,
	}
	created, err := u.userProfileRepo.CreateUserProfile(m)
	if err != nil {
		return dto.UserProfileResponseDTO{}, errors.New("Failed to create user profile: " + err.Error())
	}
	return dto.UserProfileResponseFromModel(created), nil
}

func (u UserProfileService) GetAllUserProfile() ([]dto.UserProfileResponseDTO, error) {
	list, err := u.userProfileRepo.GetAllUserProfiles()
	if err != nil {
		return []dto.UserProfileResponseDTO{}, err
	}
	out := make([]dto.UserProfileResponseDTO, 0, len(list))
	for _, x := range list {
		out = append(out, dto.UserProfileResponseFromModel(x))
	}
	return out, nil
}

func (u UserProfileService) GetUserProfileById(id int) (dto.UserProfileResponseDTO, error) {
	x, err := u.userProfileRepo.GetUserProfileById(id)
	if err != nil {
		return dto.UserProfileResponseDTO{}, errors.New("User profile not found !")
	}
	return dto.UserProfileResponseFromModel(x), nil
}

func (u UserProfileService) UpdateUserProfileEntity(newData dto.UpdateUserProfileEntityDTO) (dto.UserProfileResponseDTO, error) {
	if newData.First_name != "" && len(newData.First_name) < 4 {
		return dto.UserProfileResponseDTO{}, errors.New("First name minimum 4 characters")
	}
	if newData.Last_name != "" && len(newData.Last_name) < 4 {
		return dto.UserProfileResponseDTO{}, errors.New("Last name minimum 4 characters")
	}
	if newData.Address != "" && len(newData.Address) < 10 {
		return dto.UserProfileResponseDTO{}, errors.New("Address minimum 10 characters")
	}
	if newData.User_avatar != "" && len(newData.User_avatar) < 10 {
		return dto.UserProfileResponseDTO{}, errors.New("User avatar minimum 10 characters")
	}

	m := models.UserProfile{
		Id: newData.Id, UserId: newData.User_id, UserAvatar: newData.User_avatar,
		FirstName: newData.First_name, LastName: newData.Last_name, Address: newData.Address,
	}
	updated, err := u.userProfileRepo.UpdateUserProfile(m)
	if err != nil {
		return dto.UserProfileResponseDTO{}, err
	}
	return dto.UserProfileResponseFromModel(updated), nil
}

func (u UserProfileService) DeleteUserProfile(id int) (dto.UserProfileResponseDTO, error) {
	x, err := u.userProfileRepo.GetUserProfileById(id)
	if err != nil {
		return dto.UserProfileResponseDTO{}, errors.New("User profile not found !")
	}
	err = u.userProfileRepo.DeleteUserProfile(id)
	if err != nil {
		return dto.UserProfileResponseDTO{}, errors.New("Failed to delete user profile: " + err.Error())
	}
	return dto.UserProfileResponseFromModel(x), nil
}
