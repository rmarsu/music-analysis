package app

import (
	"context"
	"music-service/internal/server"
	handler "music-service/internal/transport/http"
	"music-service/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	handlers := handler.New()

	srv := server.New(handlers.InitRoutes())

	go func() {
		if err := srv.Run(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
