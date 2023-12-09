package routes

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"url-shortener/internal/database"
)

func GetDestAndRedirect(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	alias := chi.URLParam(r, "alias")

	//FIXME придумать что-то с фавиконом
	if alias == "favicon.ico" {
		http.Error(w, "No way", 404)
		return
	}
	var link database.Link

	result := db.Select("dest").Where("alias = ?", alias).First(&link)

	if result.Error != nil {
		//TODO use slog
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Something went wrong"))
		return
	}

	http.Redirect(w, r, link.Dest, http.StatusFound)
}
