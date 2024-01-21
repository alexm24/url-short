package main

import (
	"fmt"
	"log/slog"
	"os"

	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// init configs
	configPath := "configs/local.yaml"
	cfg := config.MustLoad(configPath)

	// init logger
	log := setupLogger(cfg.Env)

	log.Info("starting", slog.String("env", cfg.Env))

	// init storage
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		fmt.Println(err)
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init router

	// TODO: run server

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
