package handler

import (
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"github.com/gin-gonic/gin"
	"html/template"
	"math"
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
	Limit = 3
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
	offset := id
	offset -= 1
	if offset != 0 || offset != total {
		offset *= Limit
	}

	projects, err := h.service.Projects.GetProjects(Limit, offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagination := Pagination{}
	pagination.Next = id + 1
	pagination.Previous = id - 1
	pagination.CurrentPage = id
	pagination.TotalPage = int(math.Ceil(float64(total) / float64(Limit)))
	pagination.RecordPerPage = projects

	tmpl, err := template.ParseFiles("web/templates/main.html")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create templates"})
		return
	}

	c.Header("Content-Type", "text/html")
	err = tmpl.Execute(c.Writer, pagination)
	if err != nil {
		if _, err := c.Writer.WriteString("Page not found 404"); err != nil {
			return
		}
		return
	}
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
