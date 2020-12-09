package main

import (
	"context"
	"fmt"
	"github.com/golang/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serve(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()
	return http.ListenAndServe(addr, handler)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go startServer(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signalChan:
		cancel()
	}
}

func startServer(parentCtx context.Context) {
	g, ctx := errgroup.WithContext(parentCtx)

	g.Go(func() error {
		return serve(ctx, ":8080", nil)
	})

	g.Go(func() error {
		return serve(ctx, ":8081", nil)
	})
	if err := g.Wait(); err != nil {
		fmt.Println("err happen")
	}
}
