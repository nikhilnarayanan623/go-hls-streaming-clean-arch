package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/docs"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api/routes"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServerHTTP(cfg config.Config,
	videoHandler interfaces.VideoHandler, audioHandler interfaces.AudioHandler) *Server {

	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.SetupVideo(engine.Group("/video"), videoHandler)
	routes.SetupAudio(engine.Group("/audio"), audioHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (c *Server) Start() {
	c.engine.Run(c.port)
}
