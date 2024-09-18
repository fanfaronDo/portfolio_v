package service

import (
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"github.com/fanfaronDo/portfolio_v/internal/repository"
)

type Projects interface {
	GetAll() ([]domain.Project, error)
	GetTotal() (int, error)
	GetProjects(limit, offset int) ([]domain.Project, error)
}

type Project interface {
	Create(project domain.Project) error
	Delete(project_id int) error
	GetById(project_id int) (domain.Project, error)
	Update(project_id int, project domain.Project) error
}

type Service struct {
	Project
	Projects
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Project:  NewProjectService(repos),
		Projects: NewProjectsService(repos),
	}
}
