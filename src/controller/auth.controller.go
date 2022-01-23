package controller

import "github.com/gin-gonic/gin"

type IAuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type AuthController struct {
}

func (auth AuthController) Login(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "login",
	})
}

func (auth AuthController) Register(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "register",
	})
}

func NewAuthController() IAuthController {
	return &AuthController{}
}
