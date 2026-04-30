package service

import (
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/beeploop/footick/internal/db/repositories"
	"github.com/beeploop/footick/internal/model"
	"github.com/beeploop/footick/internal/utils"
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

func (t *Tracker) ContinueTask(taskID int) (model.Task, error) {
	task, err := t.Tasks.FindByID(taskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, fmt.Errorf("task with specified ID not found")
		}

		return model.Task{}, err
	}

	// Skip if already active
	if task.Status == model.TASK_ACTIVE {
		return task, fmt.Errorf("task is already active")
	}

	if task.Status == model.TASK_COMPLETED {
		return task, fmt.Errorf("cannot continue tracking an already completed task")
	}

	updatedTask, err := t.Tasks.Update(task.ID, model.UpdateTask{Status: &model.TASK_ACTIVE})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return model.Task{}, fmt.Errorf("active task already exists, pause or stop current active task to start a new one")
		}

		return model.Task{}, err
	}

	if _, err := t.Sessions.Create(model.NewSession{
		TaskID:    task.ID,
		StartedAt: time.Now(),
	}); err != nil {
		return task, err
	}

	return updatedTask, nil
}

func (t *Tracker) StopTask() (model.Task, error) {
	activeTask, err := t.Tasks.FindActive()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, fmt.Errorf("no active task to stop")
		}

		return model.Task{}, err
	}

	updatedTask, err := t.Tasks.Update(activeTask.ID, model.UpdateTask{
		Status: &model.TASK_COMPLETED,
	})
	if err != nil {
		return model.Task{}, err
	}

	session, err := t.Sessions.FindActiveByTask(updatedTask.ID)
	if err != nil {
		return model.Task{}, err
	}

	now := time.Now()
	if _, err := t.Sessions.Update(session.ID, model.UpdateSession{EndedAt: &now}); err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}

func (t *Tracker) ListTask() ([]model.TaskSession, error) {
	taskSessions := make([]model.TaskSession, 0)

	tasks, err := t.Tasks.FindWhere("status IS NOT 'completed'")
	if err != nil {
		return taskSessions, err
	}

	for _, task := range tasks {
		sessions, err := t.Sessions.FindByTaskID(task.ID)
		if err != nil {
			return taskSessions, err
		}

		taskSessions = append(taskSessions, model.TaskSession{
			Task: task,
			Sessions: slices.Collect(
				utils.Map(sessions, func(session model.Session) model.NormalizedSession {
					return model.NormalizeSession(session)
				}),
			),
		})

	}

	return taskSessions, nil
}

func (t *Tracker) SummarizeTask(taskID int, timerange model.TimeRange, completedOnly bool) ([]model.TaskSession, error) {
	taskSessions := make([]model.TaskSession, 0)

	inRangeSessions, err := t.Sessions.WithinRange(timerange.From, timerange.To)
	if err != nil {
		return taskSessions, err
	}

	uniqueTasksInSessions := utils.UniqueBy(inRangeSessions, func(session model.Session) int {
		return session.TaskID
	})

	for _, session := range uniqueTasksInSessions {
		task, err := t.Tasks.FindByID(session.TaskID)
		if err != nil {
			return taskSessions, err
		}

		if completedOnly && !task.Status.IsCompleted() {
			continue
		}

		sessions, err := t.Sessions.FindByTaskID(task.ID)
		if err != nil {
			return taskSessions, err
		}

		taskSessions = append(taskSessions, model.TaskSession{
			Task: task,
			Sessions: slices.Collect(
				utils.Map(sessions, func(session model.Session) model.NormalizedSession {
					return model.NormalizeSession(session)
				}),
			),
		})
	}

	return taskSessions, nil
}
