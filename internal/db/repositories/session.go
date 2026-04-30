package repositories

import (
	"time"

	"github.com/beeploop/footick/internal/model"
)

type SessionRepository interface {
	Create(newSession model.NewSession) (model.Session, error)
	FindByID(id int) (model.Session, error)
	FindByTaskID(taskID int) ([]model.Session, error)
	Update(id int, sessionUpdate model.UpdateSession) (model.Session, error)
	FindActiveByTask(taskID int) (model.Session, error)
	WithinRange(start time.Time, end time.Time) ([]model.Session, error)
}
