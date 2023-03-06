package service

import (
	"context"

	"github.com/sofuetakuma112/go_todo_app/entity"
	"github.com/sofuetakuma112/go_todo_app/store"
)

// storeパッケージの実装を直接参照することを避けるためにインターフェースを定義する

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	// storeパッケージのRepositoryがもつメソッドのシグネチャと一致する必要がある
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}

