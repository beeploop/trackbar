package model

import "time"

type NewTask struct {
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type Task struct {
	ID          int       `db:"id" json:"id"`
	Description string    `db:"description" json:"description"`
	IsDone      bool      `db:"is_done" json:"is_done"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
