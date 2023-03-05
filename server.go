package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

// Server構造体をインスタンス化する際にルーティング情報であるhttp.Handlerを渡す
func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// 終了シグナルを受信したときにcontextをキャンセルにする
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// errgroupを使ってゴルーチンを起動することで、
	// 立ち上げたゴルーチンで返したエラーを起動元のゴルーチンで受け取れる
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			// シャットダウン以外の要因でサーバーが停止
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// 親でcontextがキャンセルされた場合にチャネルに通知される
	<-ctx.Done()

	if err := s.srv.Shutdown(context.Background()); err != nil {
		// シャットダウン処理に失敗
		log.Printf("failed to shutdown: %+v", err)
	}

	// Goメソッドで起動した別のゴルーチンの終了を待つ（戻り値も受け取る）
	return eg.Wait()
}
