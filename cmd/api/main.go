package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/otakakot/ninshow/internal/controller"
	"github.com/otakakot/ninshow/pkg/api"
)

func main() {
	slog.Info("start ninshow server")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctr := &controller.Controller{}

	hdl, err := api.NewServer(ctr, nil)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: hdl,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	slog.Info("start server shutdown")

	ctx, cansel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cansel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	slog.Info("done server shutdown")
}
