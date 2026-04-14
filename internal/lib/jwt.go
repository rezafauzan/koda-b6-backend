package lib

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Id       int    `json:"id"`
	CartId   int    `json:"cart_id"`
	RoleName string `json:"role_name"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, cartId int, roleName string) (string, error) {
	claims := CustomClaims{
		Id:       id,
		CartId:   cartId,
		RoleName: roleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))

	if err != nil {
		return "", errors.New("Failed to login! : " + err.Error())
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("APP_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	
	return claims, nil
}
