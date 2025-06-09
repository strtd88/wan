package repo

import (
	"database/sql"
	"fmt"

	"wan/db/models"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepo(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) CreateProject(name, description, dir, status, category string) error {
	stmt := `
    INSERT INTO projects (name, description, dir, status, category)
    VALUES (?, ?, ?, ?, ?)
    `

	_, err := r.DB.Exec(stmt, name, description, dir, status, category)
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	return nil
}

func (r *ProjectRepository) GetAllProjects(filter string) ([]*models.Project, error) {
	query := "SELECT id, name, description, dir, status, category FROM projects"
	if filter != "" {
		query += " WHERE status = '" + filter + "'"
	}

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*models.Project
	for rows.Next() {
		p := &models.Project{}
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Dir, &p.Status, &p.Category)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (r *ProjectRepository) UpdateStatus(id int, status string) error {
	validStatus := map[string]bool{
		"new":      true,
		"active":   true,
		"paused":   true,
		"complete": true,
	}

	if !validStatus[status] {
		return fmt.Errorf("invalid status: %s", status)
	}

	stmt := "UPDATE projects SET status = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"
	_, err := r.DB.Exec(stmt, status, id)
	if err != nil {
		return fmt.Errorf("failed to update project status: %w", err)
	}
	return nil
}

func (r *ProjectRepository) SetObsidianNotePath(id int, path string) error {
	stmt := "UPDATE projects SET obsidian_note_path = ? WHERE id = ?"
	_, err := r.DB.Exec(stmt, path, id)
	return err
}

func (r *ProjectRepository) SetBusinessValueAchieved(id int, value string) error {
	stmt := "UPDATE projects SET business_value_achieved = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"
	_, err := r.DB.Exec(stmt, value, id)
	return err
}
