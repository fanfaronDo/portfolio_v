package handler

import (
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Pagination struct {
	Next          int
	Previous      int
	CurrentPage   int
	TotalPage     int
	RecordPerPage []domain.Project
}

const (
	Limit = 4
)

func (h *Handler) getProjects(c *gin.Context) {
	paramID, ok := c.GetQuery("projects")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": paramID + " is required"})
	}
	if paramID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID not found"})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID must be an integer"})
	}
	total, err := h.service.Projects.GetTotal()
	projects, err := h.service.Projects.GetProjects(Limit, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	pagination := Pagination{}
	pagination.Next = id + 1
	pagination.Previous = id - 1
	pagination.CurrentPage = id
	pagination.TotalPage = total
	pagination.RecordPerPage = projects

	c.JSON(http.StatusOK, pagination)
}
