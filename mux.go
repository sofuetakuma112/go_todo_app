package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/sofuetakuma112/go_todo_app/clock"
	"github.com/sofuetakuma112/go_todo_app/config"
	"github.com/sofuetakuma112/go_todo_app/handler"
	"github.com/sofuetakuma112/go_todo_app/service"
	"github.com/sofuetakuma112/go_todo_app/store"
)

// handlerはserviceに依存する?
// serviceはstoreに依存する?

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	// storeの初期化
	r := store.Repository{Clocker: clock.RealClocker{}}
	v := validator.New()

	at := &handler.AddTask{Service: &service.AddTask{DB: db, Repo: &r}, Validator: v} // ハンドラーのインスタンス化
	mux.Post("/tasks", at.ServeHTTP)                                                  // ハンドラ関数を渡す

	lt := &handler.ListTask{Service: &service.ListTask{DB: db, Repo: &r}}
	mux.Get("/tasks", lt.ServeHTTP)

	ru := &handler.RegisterUser{
		Service:   &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)

	return mux, cleanup, err
}
