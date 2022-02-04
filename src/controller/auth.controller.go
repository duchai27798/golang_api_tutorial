package controller

import (
	"github.com/duchai27798/golang_api_tutorial/src/data/dto"
	"github.com/duchai27798/golang_api_tutorial/src/data/entity"
	"github.com/duchai27798/golang_api_tutorial/src/helper"
	"github.com/duchai27798/golang_api_tutorial/src/service"
	"github.com/duchai27798/golang_api_tutorial/src/utils"
	"github.com/gin-gonic/gin"
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
	dtoOk, errDTO := helper.Validate(context.ShouldBind(&loginDTO))
	ok, errValidation := helper.Validate(loginDTO)

	if !dtoOk {
		helper.BadRequest(context, "", errDTO)
		return
	}

	if !ok {
		helper.BadRequest(context, "", errValidation)
		return
	}

	authResult := authController.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		utils.LogObj(v, "user info")
		v.Token = authController.jwtService.GenerateToken(v.ID)
		helper.Ok(context, v)
		return
	}
	helper.BadRequest(context, "Please check again your credential", helper.NewSingleError("Invalid credential"))
}

func (authController AuthController) Register(context *gin.Context) {
	var registerDTO dto.RegisterDTO
	dtoOk, errDTO := helper.Validate(context.ShouldBind(&registerDTO))
	utils.LogObj(registerDTO)
	ok, errValidation := helper.Validate(registerDTO)

	if !dtoOk {
		helper.BadRequest(context, "", errDTO)
		return
	}

	if !ok {
		helper.BadRequest(context, "", errValidation)
		return
	}

	if authController.authService.IsDuplicateEmail(registerDTO.Email) {
		helper.BadRequest(context, "", helper.NewSingleError("Duplicate Email"))
		return
	} else {
		createUser := authController.authService.CreateUser(registerDTO)
		token := authController.jwtService.GenerateToken(createUser.ID)
		createUser.Token = token
		helper.Ok(context, createUser)
		return
	}
}

func NewAuthController(authService service.IAuthService, jwtService service.IJWTService) IAuthController {
	return &AuthController{
		authService,
		jwtService,
	}
}
