package middleware

import (
	"fmt"
	"github.com/duchai27798/golang_api_tutorial/src/helper"
	"github.com/duchai27798/golang_api_tutorial/src/service"
	"github.com/duchai27798/golang_api_tutorial/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AuthorizeJWTMiddleware(jwtService service.IJWTService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.GetHeader("Authorization")
		utils.Log(authorization)
		if authorization == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authorization)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			utils.Log(fmt.Sprintf("claim[user_id]: %s", claims["user_id"]))
		} else {
			utils.LogObj(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
