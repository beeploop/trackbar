package repositories

import (
	"github.com/beeploop/footick/internal/model"
	"github.com/jmoiron/sqlx"
)

type sessionRepositoryImpl struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *sessionRepositoryImpl {
	return &sessionRepositoryImpl{
		db: db,
	}
}

func (r *sessionRepositoryImpl) Create(newSession model.NewSession) (model.Session, error) {
	var session model.Session

	result, err := r.db.NamedExec("INSERT INTO sessions (task_id, started_at) VALUES (:task_id, :started_at)", newSession)
	if err != nil {
		return session, err
	}

	sessionID, err := result.LastInsertId()
	if err != nil {
		return session, err
	}

	return r.FindByID(int(sessionID))
}

func (r *sessionRepositoryImpl) FindByID(id int) (model.Session, error) {
	var session model.Session
	if err := r.db.Get(&session, "SELECT * FROM sessions WHERE id = ?", id); err != nil {
		return session, err
	}

	return session, nil
}
