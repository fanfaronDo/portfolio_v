package handler

import (
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
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
	paramID, ok := c.GetQuery("path")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": paramID + " is required"})
		return
	}
	if paramID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID not found"})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID must be an integer"})
		return
	}

	total, err := h.service.Projects.GetTotal()
	projects, err := h.service.Projects.GetProjects(Limit, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagination := Pagination{}
	pagination.Next = id + 1
	pagination.Previous = id - 1
	pagination.CurrentPage = id
	pagination.TotalPage = total
	pagination.RecordPerPage = projects

	c.JSON(http.StatusOK, pagination.RecordPerPage)
}

func (h *Handler) getProject(c *gin.Context) {
	projectId := c.Param("id")
	id, err := strconv.Atoi(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID must be an integer"})
		return
	}

	project, err := h.service.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *Handler) createProject(c *gin.Context) {
	var project domain.Project
	if err := c.Bind(&project); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Create(project)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *Handler) updateProject(c *gin.Context) {
	projectId := c.Param("id")
	id, err := strconv.Atoi(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID must be an integer"})
		return
	}

	var project domain.Project
	if err = c.Bind(&project); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Update(id, project)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func setTemplatePath(args ...string) string {
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir)
	for _, arg := range args {
		filePath += string(filepath.Separator) + arg
	}

	return filePath
}
