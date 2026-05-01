package db

import (
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func DBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".trackbar.db"), nil
}

func Open() (*sqlx.DB, error) {
	dbPath, err := DBPath()
	if err != nil {
		return nil, err
	}

	return sqlx.Open("sqlite3", dbPath)
}
