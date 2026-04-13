package repositories

import "github.com/beeploop/footick/internal/model"

type TaskRepository interface {
	Create(newTask model.NewTask) (model.Task, error)
	FindByID(id int) (model.Task, error)
}
