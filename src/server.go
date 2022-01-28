package main

import (
	"github.com/duchai27798/golang_api_tutorial/src/config"
	"github.com/duchai27798/golang_api_tutorial/src/controller"
	"github.com/duchai27798/golang_api_tutorial/src/data/repository"
	"github.com/duchai27798/golang_api_tutorial/src/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	var (
		db             *gorm.DB                   = config.SetupDatabaseConnection()
		userRepository repository.IUserRepository = repository.NewUserRepository(db)
		authService    service.IAuthService       = service.NewAuthService(userRepository)
		jwtService     service.IJWTService        = service.NewJWTService()
		authController controller.IAuthController = controller.NewAuthController(authService, jwtService)
	)

	authRouters := router.Group("api/auth")
	{
		authRouters.POST("/login", authController.Login)
		authRouters.POST("/register", authController.Register)
	}

	router.Run(":3000")
}
