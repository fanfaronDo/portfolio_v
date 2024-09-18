package service

import (
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"github.com/fanfaronDo/portfolio_v/internal/repository"
)

type ProjectService struct {
	repo *repository.Repository
}

func NewProjectService(repo *repository.Repository) *ProjectService {
	return &ProjectService{repo}
}

func (p *ProjectService) Create(project domain.Project) error {
	return p.repo.Create(project)
}

func (p *ProjectService) Update(project_id int, project domain.Project) error {
	return p.repo.Update(project_id, project)
}

func (p *ProjectService) GetById(project_id int) (domain.Project, error) {
	return p.repo.GetById(project_id)
}

func (p *ProjectService) Delete(project_id int) error {
	return p.repo.Delete(project_id)
}
