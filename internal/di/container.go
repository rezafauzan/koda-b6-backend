package di

import (
	"rezafauzan/koda-b6-golang/internal/handlers"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/repository"
	"rezafauzan/koda-b6-golang/internal/services"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	db *pgx.Conn
	UserHandler *handlers.UserHandler
	RoleHandler *handlers.RoleHandler
}

func NewContainer() (*Container, error){
	db, err := lib.DatabaseConnect()
	if err != nil {
		return nil, err
	}
	container := &Container{
		db: db,
	}
	container.initDependencies()
	return container, nil
}

func (c *Container) initDependencies(){
	userRepo , _ := repository.NewUserRepository(c.db)
	userService := services.NewUserService(userRepo)
	c.UserHandler = handlers.NewUserHandler(userService)

	roleRepo , _ := repository.NewRoleRepository(c.db)
	roleService := services.NewRoleService(roleRepo)
	c.RoleHandler = handlers.NewRoleHandler(roleService)
}