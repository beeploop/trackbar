package repositories

import (
	"strconv"

	"github.com/Masterminds/squirrel"
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

	query, args, err := squirrel.Insert("tasks").
		Columns("description", "created_at").
		Values(newTask.Description, newTask.CreatedAt).
		ToSql()
	if err != nil {
		return task, err
	}

	result, err := r.db.Exec(query, args...)
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

	query, args, err := squirrel.Select("*").
		From("tasks").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return task, err
	}

	if err := r.db.Get(&task, query, args...); err != nil {
		return task, err
	}

	return task, nil
}

func (r *taskRepositoryImpl) Update(taskID int, taskUpdate model.UpdateTask) (model.Task, error) {
	var task model.Task

	updates := map[string]any{}

	if taskUpdate.Description != nil {
		updates["description"] = taskUpdate.Description
	}

	if taskUpdate.Status != nil {
		updates["status"] = taskUpdate.Status
	}

	query, args, err := squirrel.Update("tasks").
		SetMap(updates).
		Where(squirrel.Eq{"id": strconv.Itoa(taskID)}).
		ToSql()
	if err != nil {
		return task, err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return task, err
	}

	return r.FindByID(taskID)
}

func (r *taskRepositoryImpl) FindActive() (model.Task, error) {
	var task model.Task

	query, args, err := squirrel.Select("*").
		From("tasks").
		Where(squirrel.Eq{"status": "active"}).
		Limit(1).
		ToSql()
	if err != nil {
		return task, err
	}

	if err := r.db.Get(&task, query, args...); err != nil {
		return model.Task{}, err
	}

	return task, nil
}
