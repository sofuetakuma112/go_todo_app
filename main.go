package main

import (
	"context"
	"fmt"
	"github.com/sofuetakuma112/go_todo_app/config"
	"log"
	"net"
	"os"
)

// シャットダウン以外の理由でサーバーが停止した場合にエラーが返る
// l: 動的に選択したポート番号のリッスンを開始したnet.Listenerインターフェースを満たす型
func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mux := NewMux()
	s := NewServer(l, mux)

	return s.Run(ctx)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
