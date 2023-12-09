package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/database"
	"url-shortener/internal/logger"
	"url-shortener/src/routes"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env)
	log.Info("Config and Logger was initialized", slog.String("env", cfg.Env))
	db := database.New(cfg.Database)
	log.Info("Database connection was initialized")

	if cfg.Env == "local" {
		_ = db.AutoMigrate(&database.Link{})
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/{alias}", func(w http.ResponseWriter, r *http.Request) {
		routes.GetDestAndRedirect(w, r, db)
	})
	router.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		routes.CreateAlias(w, r, db)
	})
	router.Get("/", root)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	err := http.ListenAndServe(cfg.Server.Port, router)
	if err != nil {
		log.Error("Error on server start:", err)
		os.Exit(1)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}
