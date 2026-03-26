package di

import (
	"rezafauzan/koda-b6-golang/internal/handlers"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/repository"
	"rezafauzan/koda-b6-golang/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db                    *pgxpool.Pool
	UserHandler           *handlers.UserHandler
	UserProfileHandler    *handlers.UserProfileHandler
	UserCredentialHandler *handlers.UserCredentialHandler
	AuthHandler           *handlers.AuthHandler
	RoleHandler           *handlers.RoleHandler
	ForgotPasswordHandler *handlers.ForgotPasswordHandler
}

func NewContainer() (*Container, error) {
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

func (c *Container) initDependencies() {
	userRepo, _ := repository.NewUserRepository(c.db)
	userService := services.NewUserService(userRepo)
	c.UserHandler = handlers.NewUserHandler(userService)
	
	userProfileRepo, _ := repository.NewUserRepository(c.db)
	userProfileService := services.NewUserService(userProfileRepo)
	c.UserHandler = handlers.NewUserHandler(userProfileService)

	userCredentialRepo, _ := repository.NewUserRepository(c.db)
	userCredentialService := services.NewUserService(userCredentialRepo)
	c.UserHandler = handlers.NewUserHandler(userCredentialService)

	authService := services.NewAuthService(userRepo)
	c.AuthHandler = handlers.NewAuthHandler(authService)

	roleRepo, _ := repository.NewRoleRepository(c.db)
	roleService := services.NewRoleService(roleRepo)
	c.RoleHandler = handlers.NewRoleHandler(roleService)

	forgotPasswordRepo, _ := repository.NewForgotPasswordRepository(c.db)
	forgotPasswordService := services.NewForgotPasswordService(forgotPasswordRepo, userRepo)
	c.ForgotPasswordHandler = handlers.NewForgotPasswordHandler(forgotPasswordService)
}
