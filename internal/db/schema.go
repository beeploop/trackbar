package db

import (
	"github.com/jmoiron/sqlx"
)

func InitializeSchema(db *sqlx.DB) error {
	stmt := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		is_done BOOLEAN DEFAULT False,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		started_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ended_at DATETIME,
		FOREIGN KEY (task_id) REFERENCES tasks(id)
	);
	`

	_, err := db.Exec(stmt)
	return err
}
