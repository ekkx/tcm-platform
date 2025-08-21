package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	deps, cleanup, err := initDependencies()
	if err != nil {
		return err
	}
	defer cleanup()

	mux := initMux(deps)

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", deps.Cfg.Server.Host, deps.Cfg.Server.Port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	slog.Info("🚀 API server starting",
		slog.String("addr", srv.Addr),
		slog.String("env", string(deps.Cfg.Env)),
	)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("http server error", "err", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	slog.Info("server shutting down gracefully...")
	return srv.Shutdown(ctx)
}
