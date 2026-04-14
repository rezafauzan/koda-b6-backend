package main

import (
	"fmt"
	"os"
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/middleware"
	"rezafauzan/koda-b6-golang/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "rezafauzan/koda-b6-golang/docs"
)

// @title                       CoffeeShop
// @version                     1.0
// @description                 CoffeShop Backend Restful API
// @host                        localhost:8888
// @BasePath                    /
// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Type "Bearer" followed by a space and JWT value.
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
	routers.NewUserCredentialRouters(router, container)
	routers.NewRoleRouters(router, container)
	routers.NewForgotPasswordRouters(router, container)
	routers.NewProductRouter(router, container)
	routers.NewCartItemRouters(router, container)
	routers.NewOrderRouters(router, container)
	routers.NewProductReviewRouters(router, container)
	router.Static("/assets", "./assets")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
