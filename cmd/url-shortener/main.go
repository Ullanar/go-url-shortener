package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"os"
	"time"
	"url-shortener/cmd/url-shortener/routes"
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

	if cfg.Env == "local" {
		_, err := db.Query(
			`CREATE TABLE IF NOT EXISTS links (
    					id serial PRIMARY KEY,
    					dest VARCHAR(512),
    					alias VARCHAR(255) UNIQUE
    				);
		`)
		if err != nil {
			log.Error("DB error", err)
		}
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/{alias}", routes.GetDestAndRedirect)
	router.Post("/create", routes.CreateAlias)

	err := http.ListenAndServe(cfg.Server.Port, router)
	if err != nil {
		log.Error("Error on server start:", err)
		os.Exit(1)
	}
}
