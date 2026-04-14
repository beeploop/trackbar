package model

import "time"

type NormalizedSession struct {
	ID        int       `json:"id"`
	TaskID    int       `json:"task_id"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

type TaskSession struct {
	Task     Task                `json:"task"`
	Sessions []NormalizedSession `json:"sessions"`
}

func NormalizeSession(session Session) NormalizedSession {
	return NormalizedSession{
		ID:        session.ID,
		TaskID:    session.TaskID,
		StartedAt: session.StartedAt,
		EndedAt:   session.EndedAt.Time,
	}
}
