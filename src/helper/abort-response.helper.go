package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BadRequest api return bad request response
func BadRequest(context *gin.Context, message string, error string) {
	if message == "" {
		message = "Failed to precess request"
	}
	context.AbortWithStatusJSON(http.StatusBadRequest, BuildErrorResponse(message, error, EmptyObj{}))
}

// Ok api return response successfully
func Ok(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, BuildResponse(true, "ok", data))
}
