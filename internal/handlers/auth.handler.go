package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login godoc
// @Summary      Authenticate user
// @Description  Validates credentials and returns a JWT access token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      dto.LoginRequestDTO  true  "Login credentials"
// @Success      200   {object}  dto.Response{data=dto.LoginResponseDTO}
// @Failure      400   {object}  dto.Response
// @Failure      401   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /auth/login [post]
func (u AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequestDTO
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	result, err := u.authService.Login(req)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		switch {
		case strings.Contains(msg, "Invalid email format"):
			status = http.StatusBadRequest
		case strings.Contains(msg, "Invalid email or password"):
			status = http.StatusUnauthorized
		case strings.Contains(msg, "Failed to get user credentials by email"):
			status = http.StatusUnauthorized
		case strings.Contains(msg, "Failed to get user by email"):
			status = http.StatusUnauthorized
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: msg,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Login Success!",
		Data:    result,
	})
}

// Register godoc
// @Summary      Register user
// @Description  Creates a new user account and returns authentication result.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateUserDTO  true  "Register payload"
// @Success      201   {object}  dto.Response{data=dto.CreateUserDTO}
// @Failure      400   {object}  dto.Response
// @Failure      409   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /auth/register [post]
func (u AuthHandler) Register(ctx *gin.Context) {
	var newUser dto.CreateUserDTO

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	result, err := u.authService.Register(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Registration fail! " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Register Success!",
		Data:    result,
	})
}