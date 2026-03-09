package main

import (
	"fmt"
	"os"
	"rezafauzan/koda-b6-golang/internal/middleware"
	"rezafauzan/koda-b6-golang/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	routers.NewUserRouters(router)

	router.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
