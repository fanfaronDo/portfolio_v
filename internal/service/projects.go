package service

import (
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"github.com/fanfaronDo/portfolio_v/internal/repository"
)

type ProjectsService struct {
	repo *repository.Repository
}

func NewProjectsService(repo *repository.Repository) *ProjectsService {
	return &ProjectsService{repo: repo}
}

func (p *ProjectsService) GetProjects(limit, offset int) ([]domain.Project, error) {
	return p.repo.Projects.GetProjects(limit, offset)
}

func (p *ProjectsService) GetAll() ([]domain.Project, error) {
	return p.repo.GetAll()
}

func (p *ProjectsService) GetTotal() (int, error) {
	return p.repo.GetTotal()
}
