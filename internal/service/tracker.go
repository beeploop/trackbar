package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/beeploop/footick/internal/db/repositories"
	"github.com/beeploop/footick/internal/model"
)

type Tracker struct {
	Tasks    repositories.TaskRepository
	Sessions repositories.SessionRepository
}

func NewTrackerService(tasks repositories.TaskRepository, sessions repositories.SessionRepository) *Tracker {
	return &Tracker{
		Tasks:    tasks,
		Sessions: sessions,
	}
}

func (t *Tracker) CreateTask(description string) (model.Task, error) {
	if description == "" {
		return model.Task{}, fmt.Errorf("description is required")
	}

	// This is expected to have an error. has error && error is ErrNoRows = no active task
	_, err := t.Tasks.FindActive()
	// Means a different error occurred
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Task{}, err
	}

	// Means there is no error. An active task is found
	if err == nil {
		return model.Task{}, fmt.Errorf("active task already exists, pause or stop current active task to start a new one")
	}

	task, err := t.Tasks.Create(model.NewTask{
		Description: description,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return model.Task{}, err
	}

	if _, err := t.Sessions.Create(model.NewSession{
		TaskID:    task.ID,
		StartedAt: time.Now(),
	}); err != nil {
		return task, err
	}

	return task, nil
}

func (t *Tracker) PauseTask() (model.Task, error) {
	task, err := t.Tasks.FindActive()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, fmt.Errorf("no active task to pause")
		}
		return model.Task{}, err
	}

	updatedTask, err := t.Tasks.Update(task.ID, model.UpdateTask{Status: &model.TASK_PAUSED})
	if err != nil {
		return model.Task{}, err
	}

	session, err := t.Sessions.FindActiveByTask(task.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, fmt.Errorf("no active session with specified task ID")
		}

		return model.Task{}, err
	}

	now := time.Now()
	if _, err := t.Sessions.Update(session.ID, model.UpdateSession{EndedAt: &now}); err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}
