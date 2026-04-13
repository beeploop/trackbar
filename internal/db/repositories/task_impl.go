package repositories

import (
	"github.com/beeploop/footick/internal/model"
	"github.com/jmoiron/sqlx"
)

type taskRepositoryImpl struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *taskRepositoryImpl {
	return &taskRepositoryImpl{
		db: db,
	}
}

func (r *taskRepositoryImpl) Create(newTask model.NewTask) (model.Task, error) {
	var task model.Task

	result, err := r.db.NamedExec("INSERT INTO tasks (description, created_at) VALUES (:description, :created_at)", newTask)
	if err != nil {
		return task, err
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return task, err
	}

	return r.FindByID(int(taskID))
}

func (r *taskRepositoryImpl) FindByID(id int) (model.Task, error) {
	var task model.Task
	if err := r.db.Get(&task, "SELECT * FROM tasks WHERE id = ?", id); err != nil {
		return task, err
	}

	return task, nil
}
