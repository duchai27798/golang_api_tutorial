package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Log(data interface{}) {
	j, _ := json.MarshalIndent(data, "", "	")
	fmt.Fprintln(gin.DefaultWriter, string(j))
}

func LogError(data interface{}) {
	j, _ := json.MarshalIndent(data, "", "	")
	fmt.Fprintln(gin.DefaultWriter, string(j))
}
