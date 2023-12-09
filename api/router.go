package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"time"
	"url-shortener/api/routes"
)

func New(db *gorm.DB) *chi.Mux {
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

	return router
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_ = tmpl.Execute(w, nil)
}
