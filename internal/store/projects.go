package store

import (
	"database/sql"
	"fmt"
)

type project struct {
	Name      string
	Path      string
	CreatedAt string
}

func ListProjects(db *sql.DB) ([]project, error) {

	rows, err := db.Query("SELECT name,path,created_at FROM projects ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	defer rows.Close()

	var projects []project

	for rows.Next() {
		var project project

		if err := rows.Scan(
			&project.Name,
			&project.Path,
			&project.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read projects: %w", err)
	}

	return projects, nil
}

func AddProject(db *sql.DB, name, path string) error {
	_, err := db.Exec("INSERT INTO projects (name,path) VALUES (?,?)", name, path)

	if err != nil {
		return fmt.Errorf("failed to add project: %w", err)
	}

	return nil
}
