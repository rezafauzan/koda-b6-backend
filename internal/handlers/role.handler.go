package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (u RoleHandler) GetAllRoles(ctx *gin.Context) {
	roles, err := u.roleService.GetAllRole()
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Failed to create response get all roles! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "GET all roles",
		Results:  roles,
	})
}

func (u RoleHandler) CreateNewRole(ctx *gin.Context) {
	var newRole dto.CreateRoleDTO
	err := ctx.ShouldBind(&newRole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}
	newRoleResult, err := u.roleService.CreateNewRole(newRole)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Create Role Success !",
		Results:  newRoleResult,
	})
}

func (u RoleHandler) UpdateRole(ctx *gin.Context) {
	var newRole dto.UpdateRoleDTO
	ctx.ShouldBind(&newRole)
	updatedRole, err := u.roleService.UpdateRole(newRole)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Update Roles Success !",
		Results:  updatedRole,
	})
}

func (u RoleHandler) DeleteRole(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: "Invalid role id !",
			Results:  nil,
		})
		return
	}
	deletedRole, err := u.roleService.DeleteRole(id)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Delete role failed : " + err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Delete Roles Success !",
		Results:  deletedRole,
	})
}
