package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Log(data string) {
	j, _ := json.MarshalIndent(data, "", "	")
	fmt.Fprintln(gin.DefaultWriter, string(j))
}

func LogObj(data interface{}, params ...interface{}) {
	j, _ := json.MarshalIndent(data, "", "	")
	if len(params) > 0 {
		fmt.Fprintln(gin.DefaultWriter, params[0])
	}
	fmt.Fprintln(gin.DefaultWriter, string(j))
}

func LogError(data interface{}) {
	LogObj(data)
	//panic(data)
}
