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
	router.Use(h.middlewareCORS())
	projects := router.Group("/projects")
	{
		projects.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
			return
		})

		projects.GET("", h.getProjects)
		projects.GET("/:id", h.getProject)

	}

	admin := router.Group("/admin")
	{
		admin.POST("", h.createProject)
		admin.PATCH("", h.updateProject)
	}

	return router
}
