package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api/handler/interfaces"
)

func SetupVideo(api *gin.RouterGroup, videoHandler interfaces.VideoHandler) {

	api.POST("/", videoHandler.Upload)
	api.GET("/all", videoHandler.FindAll)
	api.GET("/stream/:video_id/:playlist", videoHandler.Stream)
}
