package handler

import (
	"context"

	"github.com/sofuetakuma112/go_todo_app/entity"
)

// serviceパッケージの実装を直接参照することを避けるためにインターフェースを定義する

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService
type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}
