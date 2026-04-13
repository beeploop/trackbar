package app

import (
	"github.com/beeploop/footick/internal/db"
	"github.com/beeploop/footick/internal/db/repositories"
	"github.com/beeploop/footick/internal/service"
)

type App struct {
	Tracker *service.Tracker
}

func Bootstrap() (*App, error) {
	sqliteDB, err := db.Open()
	if err != nil {
		return nil, err
	}

	if err := db.InitializeSchema(sqliteDB); err != nil {
		return nil, err
	}

	taskRepository := repositories.NewTaskRepository(sqliteDB)
	sessionRepository := repositories.NewSessionRepository(sqliteDB)

	trackerService := service.NewTrackerService(taskRepository, sessionRepository)

	return &App{
		Tracker: trackerService,
	}, nil
}
