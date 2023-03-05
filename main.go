package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

// シャットダウン以外の理由でサーバーが停止した場合にエラーが返る
// l: 動的に選択したポート番号のリッスンを開始したnet.Listenerインターフェースを満たす型
func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		// Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	// errgroupを使ってゴルーチンを起動することで、
	// 立ち上げたゴルーチンで返したエラーを起動元のゴルーチンで受け取れる
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// サーバーを停止したときerrが返ってくる？
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// 親でcontextがキャンセルされた場合にチャネルに通知される
	<-ctx.Done()

	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	// Goメソッドで起動した別のゴルーチンの終了を待つ
	return eg.Wait()
}

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
