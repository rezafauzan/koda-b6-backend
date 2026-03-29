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

func (u *UserProfileService) GetUserProfileByUserId(userId int) (dto.UserProfileResponseDTO, error) {
	profile, err := u.userProfileRepo.GetUserProfileByUserId(userId)

	if err != nil {
		return dto.UserProfileResponseDTO{}, errors.New("User profile not found !")
	}

	modeledData := dto.UserProfileResponseDTO{
		Id:         profile.Id,
		UserId:     profile.UserId,
		UserAvatar: profile.UserAvatar,
		FirstName:  profile.FirstName,
		LastName:   profile.LastName,
		Address:    profile.Address,
	}

	return modeledData, nil
}

func (u UserProfileService) UpdateUserProfile(newData dto.UpdateUserProfileDTO, userId int) (dto.UserProfileResponseDTO, error) {
	userProfile, err := u.GetUserProfileByUserId(userId)
	if err != nil {
		return dto.UserProfileResponseDTO{}, err
	}

	if newData.User_avatar == "" {
		newData.User_avatar = userProfile.UserAvatar
	}

	if newData.First_name == "" {
		newData.First_name = userProfile.FirstName
	}

	if newData.Last_name == "" {
		newData.Last_name = userProfile.LastName
	}

	if newData.Address == "" {
		newData.Address = userProfile.Address
	}

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

	modeledData := models.UserProfile{
		Id:         newData.Id,
		UserId:     userId,
		UserAvatar: newData.User_avatar,
		FirstName:  newData.First_name,
		LastName:   newData.Last_name,
		Address:    newData.Address,
	}

	updated, err := u.userProfileRepo.UpdateUserProfile(modeledData)
	if err != nil {
		return dto.UserProfileResponseDTO{}, err
	}
	response := dto.UserProfileResponseDTO{
		Id:         updated.Id,
		UserId:     updated.UserId,
		UserAvatar: updated.UserAvatar,
		FirstName:  updated.FirstName,
		LastName:   updated.LastName,
		Address:    updated.Address,
		CreatedAt:  updated.CreatedAt,
		UpdatedAt:  updated.UpdatedAt,
	}
	return response, nil
}
