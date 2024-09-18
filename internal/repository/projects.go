package repository

import (
	"database/sql"
	"github.com/fanfaronDo/portfolio_v/internal/domain"
)

type ProjectsMysql struct {
	db *sql.DB
}

func NewProjects(db *sql.DB) *ProjectsMysql {
	return &ProjectsMysql{db: db}
}

func (p *ProjectsMysql) GetProjects(limit, offset int) ([]domain.Project, error) {
	var projects []domain.Project
	query := "SELECT title, description, image FROM projects OFFSET $1 LIMIT $2"
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		return projects, err
	}

	defer rows.Close()
	for rows.Next() {
		var project domain.Project
		err = rows.Scan(&project.Title, &project.Description, &project.Image)
		if err != nil {
			return projects, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (p *ProjectsMysql) GetAll() ([]domain.Project, error) {
	var projects []domain.Project
	rows, err := p.db.Query("SELECT title, description, image FROM projects")
	if err != nil {
		return projects, err
	}
	defer rows.Close()

	for rows.Next() {
		var project domain.Project
		err = rows.Scan(&project.Title, &project.Description, &project.Image)
		if err != nil {
			return projects, err
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return projects, err
	}

	return projects, nil
}

func (p *ProjectsMysql) GetTotal() (int, error) {
	var total int
	err := p.db.QueryRow("SELECT count(project_id) FROM projects").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
