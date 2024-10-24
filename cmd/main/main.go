package main

import (
	"go-url-shortener-service/internal/config"
	"go-url-shortener-service/internal/lib/slog/sl"
	"go-url-shortener-service/internal/storage/sqlite"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main (){

	//cfg
	cfg := config.MustLoad()

	//logger
	log := setupLogger(cfg.Env)
	log.Info("starting service", slog.String("env", cfg.Env))
	log.Debug("debug logger is on")


	//storage
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init db", sl.Err(err) )
		os.Exit(1)

	}

	_ = storage

	//router
	//server
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env{
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		
	}
	return log

}