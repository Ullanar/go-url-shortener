package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("Config and Logger was initialized")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "production":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
