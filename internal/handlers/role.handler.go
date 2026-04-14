package handlers

import (
	"errors"
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

// GetAllRoles godoc
// @Summary      List roles
// @Description  Returns all roles in the system.
// @Tags         roles
// @Produce      json
// @Success      200  {object}  dto.Response{data=[]dto.RoleResponseDTO}
// @Failure      500  {object}  dto.Response
// @Router       /roles [get]
func (u RoleHandler) GetAllRoles(ctx *gin.Context) {
	roles, err := u.roleService.GetAllRole()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to create response get all roles! : " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "GET all roles",
		Data:    roles,
	})
}

// CreateNewRole godoc
// @Summary      Create role
// @Description  Creates a new role with a unique name (minimum 4 characters).
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateRoleDTO  true  "Role payload"
// @Success      201   {object}  dto.Response{data=dto.RoleResponseDTO}
// @Failure      400   {object}  dto.Response
// @Failure      409   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /roles [post]
func (u RoleHandler) CreateNewRole(ctx *gin.Context) {
	var newRole dto.CreateRoleDTO
	err := ctx.ShouldBindJSON(&newRole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	newRoleResult, err := u.roleService.CreateNewRole(newRole)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		switch {
		case strings.Contains(msg, "minimum is 4"):
			status = http.StatusBadRequest
		case strings.Contains(msg, "unique") || strings.Contains(msg, "duplicate") || strings.Contains(msg, "23505") || strings.Contains(strings.ToLower(msg), "already exists"):
			status = http.StatusConflict
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: msg,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Create Role Success !",
		Data:    newRoleResult,
	})
}

// UpdateRole godoc
// @Summary      Update role
// @Description  Partially updates an existing role by id.
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        body  body      dto.UpdateRoleDTO  true  "Role update payload"
// @Success      200   {object}  dto.Response{data=dto.RoleResponseDTO}
// @Failure      400   {object}  dto.Response
// @Failure      404   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /roles [patch]
func (u RoleHandler) UpdateRole(ctx *gin.Context) {
	var newRole dto.UpdateRoleDTO
	err := ctx.ShouldBindJSON(&newRole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "id is required",
			Data:    nil,
		})
		return
	}
	newRole.Id = id
	updatedRole, err := u.roleService.UpdateRole(newRole)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		if errors.Is(err, pgx.ErrNoRows) || strings.Contains(msg, "no rows") {
			status = http.StatusNotFound
		} else if strings.Contains(msg, "minimum 4") {
			status = http.StatusBadRequest
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
		Message: "Update Roles Success !",
		Data:    updatedRole,
	})
}

// DeleteRole godoc
// @Summary      Delete role
// @Description  Deletes a role by id.
// @Tags         roles
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.Response
// @Failure      404  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /roles/{id} [delete]
func (u RoleHandler) DeleteRole(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid role id !",
			Data:    nil,
		})
		return
	}
	deleted, err := u.roleService.DeleteRole(id)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		if strings.Contains(msg, "Role not found") {
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: "Delete role failed : " + msg,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Role deleted successfully",
		Data:    deleted,
	})
}

func (u RoleHandler) GetRoleByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if strings.TrimSpace(name) == "" {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "role name is required",
			Data:    nil,
		})
		return
	}

	role, err := u.roleService.GetRoleByName(name)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(msg), "not found") {
			status = http.StatusNotFound
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
		Message: "Success get role",
		Data:    role,
	})
}
