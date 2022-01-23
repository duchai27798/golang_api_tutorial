package main

import (
	"github.com/duchai27798/golang_api_tutorial/src/controller"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var (
	//db             *gorm.DB = config.SetupDatabaseConnection()
	authController = controller.NewAuthController()
)

func main() {
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	authRouters := router.Group("api/auth")
	{
		authRouters.POST("/login", authController.Login)
		authRouters.POST("/register", authController.Register)
	}

	router.Run(":3000")
}
