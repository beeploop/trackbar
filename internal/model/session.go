package model

import (
	"database/sql"
	"time"
)

type NewSession struct {
	TaskID    int       `db:"task_id"`
	StartedAt time.Time `db:"started_at"`
}

type Session struct {
	ID        int          `db:"id" json:"id"`
	TaskID    int          `db:"task_id" json:"task_id"`
	StartedAt time.Time    `db:"started_at" json:"started_at"`
	EndedAt   sql.NullTime `db:"ended_at" json:"ended_at"`
}
