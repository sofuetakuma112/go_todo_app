package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sofuetakuma112/go_todo_app/config"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("falied to read env: %v", cfg)
	}

	sut, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		t.Fatalf("falied to create mux: %v", err)
	}

	t.Cleanup(cleanup)

	sut.ServeHTTP(w, r)
	resp := w.Result()
	// Cleanupは、テスト（またはサブテスト）と
	// そのすべてのサブテストが完了したときに呼び出される関数を登録します。
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("falied to read body: %v", err)
	}

	want := `{"status": "ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
