package repositories

import "github.com/beeploop/trackbar/internal/model"

type TaskRepository interface {
	Create(newTask model.NewTask) (model.Task, error)
	FindByID(id int) (model.Task, error)
	FindWhere(query string, args ...interface{}) ([]model.Task, error)
	Update(id int, taskUpdate model.UpdateTask) (model.Task, error)
	FindActive() (model.Task, error)
}
