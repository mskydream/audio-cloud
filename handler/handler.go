package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mskydream/audio-cloud/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refreshTokens)
	}

	main := router.Group("/main", h.userIdentity)
	{
		audio := main.Group("/audio")
		{
			audio.GET("/", h.getAllAudio)
			audio.POST("/", h.uploadAudio)
			audio.PUT("/:id", h.addDescription)
			audio.GET("/:id", h.downloadAudio)
		}

		share := main.Group("share")
		{
			share.POST("/:id", h.shareAudio)
			share.DELETE("/:id", h.unshareAudio)
		}

		main.GET("/shares", h.getSharedAudio)
	}

	return router
}
