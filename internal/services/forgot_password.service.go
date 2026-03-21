package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strconv"
	"strings"
)

type ForgotPasswordService struct {
	forgotPasswordRepo *repository.ForgotPasswordRepository
	userRepo           *repository.UserRepository
}

func NewForgotPasswordService(forgotPasswordRepo *repository.ForgotPasswordRepository, userRepo *repository.UserRepository) *ForgotPasswordService {
	return &ForgotPasswordService{
		forgotPasswordRepo: forgotPasswordRepo,
		userRepo:           userRepo,
	}
}

func (f ForgotPasswordService) RequestForgotPassword(email string) (models.ForgotPassword, error) {
	if !strings.Contains(email, "@") {
		return models.ForgotPassword{}, errors.New("Failed to request forgot password! : Invalid email format !")
	}

	user, userExists, err := f.userRepo.GetUserByEmail(email)
	if err != nil {
		return models.ForgotPassword{}, errors.New("Failed to request forgot password! : Email not found !")
	}
	if !userExists {
		return models.ForgotPassword{}, errors.New("Failed to request forgot password! : Email not found !")
	}

	codeOTP, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return models.ForgotPassword{}, errors.New("Failed to request forgot password! : " + err.Error())
	}

	otp := int(codeOTP.Int64()) + 100000

	forgotPasswordRequest, err := f.forgotPasswordRepo.CreateForgotPasswordData(user.Email, otp)
	if err != nil {
		return models.ForgotPassword{}, err
	}
	return forgotPasswordRequest, nil
}

func (f ForgotPasswordService) ResetPassword(req dto.ResetForgotPasswordDTO) error {
	if !strings.Contains(req.Email, "@") {
		return errors.New("Failed to reset password! : Invalid email format !")
	}
	if len(req.New_password) < 8 {
		return errors.New("Failed to reset password! : Password too weak minimum length is 8 characters !")
	}
	if req.New_password != req.Password_confirm {
		return errors.New("Failed to reset password! : Password confirmation missmatch !")
	}
	fmt.Println(req.Email)
	_, _, err := f.userRepo.GetUserByEmail(req.Email)

	if err != nil {
		return errors.New("Failed to reset password! : Email not found !")
	}

	latestOTP, err := f.forgotPasswordRepo.GetLatestOTP(req.Email)
	if err != nil {
		return errors.New("Failed to reset password! : OTP not found !")
	}

	otpRaw := strings.TrimSpace(req.Otp)
	if otpRaw == "" {
		otpRaw = strings.TrimSpace(req.Code_otp)
	}
	if otpRaw == "" {
		return errors.New("Failed to reset password! : OTP is invalid !")
	}

	otpAsNumber, err := strconv.Atoi(otpRaw)
	if err != nil {
		return errors.New("Failed to reset password! : OTP is invalid !")
	}
	if latestOTP.CodeOtp != otpAsNumber {
		return errors.New("Failed to reset password! : OTP is invalid !")
	}

	err = f.userRepo.UpdatePassword(req.Email, req.New_password)
	if err != nil {
		return errors.New("Failed to reset password! : " + err.Error())
	}

	_ = f.forgotPasswordRepo.MarkOTPUsed(latestOTP.Id)

	return nil
}
