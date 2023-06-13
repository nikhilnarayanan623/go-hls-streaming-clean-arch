package interfaces

import "github.com/gin-gonic/gin"

type VideoHandler interface {
	Upload(ctx *gin.Context)
}
