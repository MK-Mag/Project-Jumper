package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "turso.tech/database/tursogo"
)

func Open() (*sql.DB, error) {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return nil, fmt.Errorf("LOCALAPPDATA environment variable is not set")
	}

	dataDir := filepath.Join(localAppData, "PJ")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("create data directory: %w", err)
	}

	dbPath := filepath.Join(dataDir, "PJ.db")

	db, err := sql.Open("turso", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Open DB :%W", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("Connect ro database: %w", err)
	}

	const schema = `
		CREATE TABLE IF NOT EXISTS projects(
			id INTEGER PRIMART KEY,
			name TEXT NOT NULL UNIQUE,
			path TEXT NOT NULL,
			created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return nil, fmt.Errorf("create table: %w", err)
	}

	return nil, nil
}
