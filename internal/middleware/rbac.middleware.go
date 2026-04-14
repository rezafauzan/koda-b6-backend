package middleware

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"strings"

	"github.com/gin-gonic/gin"
)

func RBAC(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleVal, exists := ctx.Get("role_name")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "Unauthorized",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		roleName := strings.ToLower(roleVal.(string))
		for _, allowed := range allowedRoles {
			if roleName == strings.ToLower(allowed) {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: "Forbidden: Access denied",
			Data:    nil,
		})
		ctx.Abort()
	}
}
