package main

import (
	"log/slog"
	"url-shortener/internal/config"
	"url-shortener/internal/database"
	"url-shortener/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env)
	log.Info("Config and Logger was initialized", slog.String("env", cfg.Env))
	db := database.New(cfg.Database)

	if cfg.Env == "local" {
		_, err := db.Query(
			`CREATE TABLE IF NOT EXISTS links (
    					id serial PRIMARY KEY,
    					dest VARCHAR(512),
    					alias VARCHAR(255)
    				);
		`)
		if err != nil {
			log.Error("DB error", err)
		}
	}
}
