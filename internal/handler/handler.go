package handler

import (
	"github.com/fanfaronDo/portfolio_v/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	projects := router.Group("/")
	{
		projects.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
			return
		})

		projects.GET("/", h.getProjects)
	}

	return router
}
