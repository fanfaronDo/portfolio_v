package repository

import (
	"database/sql"
	"github.com/fanfaronDo/portfolio_v/internal/domain"
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

type Repository struct {
	Project
	Projects
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Project:  NewProject(db),
		Projects: NewProjects(db),
	}
}
