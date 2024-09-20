package handler

import (
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
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
	paramID, ok := c.GetQuery("projects")
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

	filePath := setTemplatePath("web", "template", "main.html")
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmp, err := template.New("projects").Parse(string(file))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = tmp.Execute(c.Writer, pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func setTemplatePath(args ...string) string {
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir)
	for _, arg := range args {
		filePath += string(filepath.Separator) + arg
	}

	return filePath
}
