package controller

import (
	"github.com/duchai27798/golang_api_tutorial/src/data/dto"
	"github.com/duchai27798/golang_api_tutorial/src/data/entity"
	"github.com/duchai27798/golang_api_tutorial/src/helper"
	"github.com/duchai27798/golang_api_tutorial/src/service"
	"github.com/duchai27798/golang_api_tutorial/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type AuthController struct {
	authService service.IAuthService
	jwtService  service.IJWTService
}

func (authController AuthController) Login(context *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := context.ShouldBind(&loginDTO)
	utils.LogObj(loginDTO)
	if errDTO != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("Failed to precess request", errDTO.Error(), helper.EmptyObj{}))
		return
	}

	authResult := authController.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		utils.LogObj(v, "user info")
		v.Token = authController.jwtService.GenerateToken(v.ID)
		context.JSON(http.StatusOK, helper.BuildResponse(true, "ok", v))
		return
	}
	context.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("Please check again your credential", "Invalid credential", helper.EmptyObj{}))
}

func (authController AuthController) Register(context *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := context.ShouldBind(&registerDTO)
	if errDTO != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("Failed to precess request", errDTO.Error(), helper.EmptyObj{}))
		return
	}

	if !authController.authService.IsDuplicateEmail(registerDTO.Email) {
		context.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse("Failed to precess request", "Duplicate Email", helper.EmptyObj{}))
		return
	} else {
		createUser := authController.authService.CreateUser(registerDTO)
		token := authController.jwtService.GenerateToken(createUser.ID)
		createUser.Token = token
		context.JSON(http.StatusOK, helper.BuildResponse(true, "ok", createUser))
		return
	}
}

func NewAuthController(authService service.IAuthService, jwtService service.IJWTService) IAuthController {
	return &AuthController{
		authService,
		jwtService,
	}
}
