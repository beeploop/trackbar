package repositories

import "github.com/beeploop/footick/internal/model"

type SessionRepository interface {
	Create(newSession model.NewSession) (model.Session, error)
	FindByID(id int) (model.Session, error)
}
