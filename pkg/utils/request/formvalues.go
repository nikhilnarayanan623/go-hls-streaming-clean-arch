package request

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFormValues(ctx *gin.Context, name string) (value string, err error) {

	value = ctx.Request.PostFormValue(name)
	if value == "" {
		return "", fmt.Errorf("failed to get %s from request", name)
	}

	return value, nil
}
