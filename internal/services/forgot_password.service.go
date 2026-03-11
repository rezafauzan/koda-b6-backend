package services

import (
	"crypto/rand"
	"math/big"
	"rezafauzan/koda-b6-golang/internal/models"
	"rezafauzan/koda-b6-golang/internal/repository"
)

type ForgotPasswordService struct {
	ForgotPasswordRepo *repository.ForgotPasswordRepository
	UserRepo           *repository.UserRepository
}

func (f ForgotPasswordService) NewForgotPasswordService() *ForgotPasswordService {
	return &ForgotPasswordService{
		ForgotPasswordRepo: &repository.ForgotPasswordRepository{},
		UserRepo:           &repository.UserRepository{},
	}
}

func (f ForgotPasswordService) RequestForgotPassword(email string) (models.ForgotPassword, error) {
	user, err := f.UserRepo.GetUserByEmail(email)
	if err != nil {
		return models.ForgotPassword{}, err
	}
	code_otp, err := rand.Int(rand.Reader, big.NewInt(900000))
	forgotPasswordRequest, err := f.ForgotPasswordRepo.CreateForgotPasswordData(user.Email, code_otp)
	if err != nil {
		return models.ForgotPassword{}, err
	}
	return forgotPasswordRequest, nil
}