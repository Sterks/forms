package app

import (
	"context"
	"errors"
	"fmt"
	"forms/internal/config"
	delivery "forms/internal/delivery/http"
	"forms/internal/repository"
	"forms/internal/server"
	"forms/internal/service"
	"forms/pkg/logger"
	"forms/pkg/mongodb"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func Run(configsPath string) {
	checkDir()

	cfg, err := config.Initial(configsPath)
	if err != nil {
		logger.Errorf("не могу подулючиться %v", err)
		return
	}

	mongoClient, err := mongodb.NewClient(cfg.Mongo.URL, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		logger.Errorf("Нет подключения %v", err)
	}

	db := mongoClient.Database(cfg.Mongo.DBName)

	repos := repository.NewRepository(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handlers := delivery.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}
}

func checkDir() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}
