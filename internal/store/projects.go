package store

import (
	"database/sql"
	"errors"
	"fmt"
)

type Project struct {
	Name      string
	Path      string
	CreatedAt string
}

func GetProject(db *sql.DB, name string) (Project, error) {

	row := db.QueryRow("SELECT name,path,created_at FROM Projects WHERE name = ?", name)

	var project Project
	err := row.Scan(&project.Name, &project.Path, &project.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return Project{}, fmt.Errorf("project %q not found", name)
	}

	if err != nil {
		return Project{}, fmt.Errorf("failed to get Project: %w", err)
	}

	return project, nil
}

func ListProjects(db *sql.DB) ([]Project, error) {

	rows, err := db.Query("SELECT name,path,created_at FROM Projects ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("failed to list Projects: %w", err)
	}
	defer rows.Close()

	var Projects []Project

	for rows.Next() {
		var Project Project

		if err := rows.Scan(
			&Project.Name,
			&Project.Path,
			&Project.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan Project: %w", err)
		}

		Projects = append(Projects, Project)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read Projects: %w", err)
	}

	return Projects, nil
}

func AddProject(db *sql.DB, name, path string) error {
	_, err := db.Exec("INSERT INTO Projects (name,path) VALUES (?,?)", name, path)

	if err != nil {
		return fmt.Errorf("failed to add Project: %w", err)
	}

	return nil
}

func RemoveProject(db *sql.DB, name string) error {
	result,err:=db.Exec("DELETE FROM projects WHERE name = ?", name)
	if err !=nil{
		return fmt.Errorf("failed to remove Project: %w", err)
	}

	rowsAffected,err :=result.RowsAffected()
	if err != nil{
		return fmt.Errorf("failed to check removed project: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("project %q not found", name)
	}

	return nil
}
