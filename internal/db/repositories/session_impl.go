package repositories

import (
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/beeploop/trackbar/internal/model"
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

	query, args, err := squirrel.Insert("sessions").
		Columns("task_id", "started_at").
		Values(newSession.TaskID, newSession.StartedAt).
		ToSql()
	if err != nil {
		return session, err
	}

	result, err := r.db.Exec(query, args...)
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

	query, args, err := squirrel.Select("*").
		From("sessions").
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		ToSql()
	if err != nil {
		return session, err
	}

	if err := r.db.Get(&session, query, args...); err != nil {
		return session, err
	}

	return session, nil
}

func (r *sessionRepositoryImpl) FindByTaskID(taskID int) ([]model.Session, error) {
	sessions := make([]model.Session, 0)

	query, args, err := squirrel.Select("*").
		From("sessions").
		Where(squirrel.Eq{"task_id": strconv.Itoa(taskID)}).
		ToSql()
	if err != nil {
		return sessions, err
	}

	if err := r.db.Select(&sessions, query, args...); err != nil {
		return sessions, err
	}

	return sessions, nil
}

func (r *sessionRepositoryImpl) Update(sessionID int, sessionUpdate model.UpdateSession) (model.Session, error) {
	var session model.Session

	updates := map[string]any{}

	if sessionUpdate.EndedAt != nil {
		updates["ended_at"] = sessionUpdate.EndedAt
	}

	query, args, err := squirrel.Update("sessions").
		SetMap(updates).
		Where(squirrel.Eq{"id": sessionID}).
		ToSql()
	if err != nil {
		return session, err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return session, err
	}

	return r.FindByID(sessionID)
}

func (r *sessionRepositoryImpl) FindActiveByTask(taskID int) (model.Session, error) {
	var session model.Session

	query, args, err := squirrel.Select("*").
		From("sessions").
		Where(squirrel.Eq{
			"task_id":  strconv.Itoa(taskID),
			"ended_at": nil,
		}).
		ToSql()
	if err != nil {
		return session, err
	}

	if err := r.db.Get(&session, query, args...); err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (r *sessionRepositoryImpl) WithinRange(start time.Time, end time.Time) ([]model.Session, error) {
	sessions := make([]model.Session, 0)

	query, args, err := squirrel.Select("*").
		From("sessions").
		Where(squirrel.Lt{"started_at": end}).
		Where(squirrel.Or{
			squirrel.Eq{"ended_at": nil},
			squirrel.Gt{"ended_at": start},
		}).ToSql()
	if err != nil {
		return sessions, err
	}

	if err := r.db.Select(&sessions, query, args...); err != nil {
		return sessions, err
	}

	return sessions, nil
}
