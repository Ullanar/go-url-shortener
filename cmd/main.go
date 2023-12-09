package main

import (
	"log/slog"
	"net/http"
	"os"
	"url-shortener/api"
	"url-shortener/internal/config"
	"url-shortener/internal/database"
	"url-shortener/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env)
	log.Info("Config and Logger was initialized", slog.String("env", cfg.Env))
	db := database.New(cfg.Database)
	log.Info("Database connection was initialized")

	router := api.New(db)

	if cfg.Env == "local" {
		_ = db.AutoMigrate(&database.Link{})
	}

	err := http.ListenAndServe(cfg.Server.Port, router)
	if err != nil {
		log.Error("Error on server start:", err)
		os.Exit(1)
	}
}
