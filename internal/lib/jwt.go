package lib

import (
	"errors"
	"os"
	"rezafauzan/koda-b6-golang/internal/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	User_id int `json:"id"`
	Role string `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(user dto.UserResponseDTO) (string, error) {
	claims := CustomClaims{
		User_id: user.Id,
		Role: user.Role_name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", errors.New("Failed to login! : " + err.Error())
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
