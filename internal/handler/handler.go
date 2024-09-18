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
	projectsRange := router.Group("/")
	{
		projectsRange.GET("/", h.getProjects)
	}

	return router
}
