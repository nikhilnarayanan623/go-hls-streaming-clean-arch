package interfaces

import "github.com/gin-gonic/gin"

type VideoHandler interface {
	Upload(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Stream(ctx *gin.Context)
}
