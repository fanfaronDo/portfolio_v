package repository

import (
	"database/sql"
	"fmt"
	"github.com/fanfaronDo/portfolio_v/internal/domain"
	"log"
)

type ProjectMysql struct {
	db *sql.DB
}

func NewProject(db *sql.DB) *ProjectMysql {
	return &ProjectMysql{
		db: db,
	}
}

func (p *ProjectMysql) Create(project domain.Project) error {
	query := "insert into projects (title, description, image) values (?, ?, ?)"
	insert, err := p.db.Query(query, project.Title, project.Description, project.Image)
	defer insert.Close()
	if err != nil {
		log.Printf("Project %s is not added: %v\n", project.Title, err)
		return err
	}

	return nil
}

func (p *ProjectMysql) Delete(project_id int) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	del, err := tx.Query("delete from projects where id=?", project_id)
	set, err := tx.Query("SET @newid := 0")
	update, err := tx.Query("UPDATE projects SET id = (@newid := @newid + 1)")

	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	if err != nil {
		return err
	}
	defer func() {
		del.Close()
		set.Close()
		update.Close()
	}()

	return nil
}

func (p ProjectMysql) GetById(project_id int) (domain.Project, error) {
	var prj domain.Project
	query := "SELECT title, description, image FROM projects WHERE id = ?"
	err := p.db.QueryRow(query, project_id).Scan(&prj.Title, &prj.Description, &prj.Image)

	if err != nil {
		if err == sql.ErrNoRows {
			return prj, fmt.Errorf("project with id %d not found", project_id)
		}
		return prj, err
	}

	return prj, nil
}

func (p ProjectMysql) Update(project_id int, project domain.Project) error {
	query := "update projects set title=$1, description=$1, image=$1 where id=?"
	update, err := p.db.Query(query, project.Title, project.Description, project.Image, project_id)
	if err != nil {
		log.Printf("Update project with id %d failed: %v\n", project, err)
		return err
	}
	defer update.Close()

	return nil
}
