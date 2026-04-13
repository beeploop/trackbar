package service

import (
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
