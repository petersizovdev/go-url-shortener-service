package main

import (
	"go-url-shortener-service/internal/config"
	"go-url-shortener-service/internal/storage/sqlite"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//cfg
	cfg := config.MustLoad()

	//logger
	log := setupLogger(cfg.Env)
	log.Info("starting service", slog.String("env", cfg.Env))
	log.Debug("debug logger is on")

	//storage
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init db", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer storage.Close()

	//router
	router := chi.NewRouter()
		//middleware
		router.Use(middleware.RequestID)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Use(middleware.URLFormat)



	//server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}