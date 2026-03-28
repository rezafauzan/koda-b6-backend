package main

import (
	"fmt"
	"os"
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/middleware"
	"rezafauzan/koda-b6-golang/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer recover()
	godotenv.Load()

	router := gin.Default()
	container, err := di.NewContainer()

	if err != nil {
		panic("Container Error : " + err.Error())
	}

	router.Use(middleware.CORSMiddleware())

	routers.NewAuthRouters(router, container)
	routers.NewUserRouters(router, container)
	routers.NewUserProfileRouters(router, container)
	routers.NewRoleRouters(router, container)
	routers.NewForgotPasswordRouters(router, container)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
