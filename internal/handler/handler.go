package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/dockerPing/internal/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	ping := router.Group("/ping")
	ping.POST("/", h.addPing)

	return router
}

