package service

import (
	"context"
	"fmt"

	"github.com/sofuetakuma112/go_todo_app/entity"
	"github.com/sofuetakuma112/go_todo_app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister // storeパッケージの実装へのアクセスはserviceパッケージで定義したインターフェースを介して行う
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
