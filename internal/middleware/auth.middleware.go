package middleware

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Messages: "Authorization header required",
				Results: nil,
			})
			ctx.Abort()
			return
		}
		
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Messages: "Invalid token format",
				Results: nil,
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := lib.VerifyJWT(tokenString)
		
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Messages: "Invalid or expired token",
				Results: nil,
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.User_id)
		ctx.Set("role_id", claims.Role)

		ctx.Next()
	}
}
