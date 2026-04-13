package model

import "time"

type TaskStatus string

var (
	TASK_ACTIVE    TaskStatus = "active"
	TASK_PAUSED    TaskStatus = "paused"
	TASK_COMPLETED TaskStatus = "completed"
)

type NewTask struct {
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type Task struct {
	ID          int        `db:"id" json:"id"`
	Description string     `db:"description" json:"description"`
	Status      TaskStatus `db:"status" json:"status"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
}

type UpdateTask struct {
	Description *string     `db:"description"`
	Status      *TaskStatus `db:"status"`
}
