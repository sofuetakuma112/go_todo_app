package store

import (
	"context"

	"github.com/sofuetakuma112/go_todo_app/entity"
)

// DBへのコネクション(*sqlx.Db, *sqlx.Txなど)を引数として受け取る

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, // *sqlx.Tx, *sqlx.DbともにQueryerを満たす
) (entity.Tasks, error) {
	tasks := entity.Tasks{} // スライスを初期化
	sql := `SELECT id, title, status, created, modified FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task (title, status, created, modified) VALUES (?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, sql, t.Title, t.Status, t.Created, t.Modified)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // INSERT時に発行されたidを取得
	if err != nil {
		return err
	}

	t.ID = entity.TaskID(id) // 引数の*entity.TaskのIDフィールドを更新して、呼び出し元に発行されたIDを伝える
	return nil
}
