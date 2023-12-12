package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"url-shortener/internal/repository"
)

func GetDestAndRedirect(w http.ResponseWriter, r *http.Request, repository repository.Repository) {

	alias := chi.URLParam(r, "alias")

	//FIXME придумать что-то с фавиконом
	if alias == "favicon.ico" {
		http.Error(w, "No way", 404)
		return
	}
	dest, err := repository.GetDestinationByAlias(alias)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Something went wrong"))
		return
	}

	http.Redirect(w, r, dest, http.StatusFound)
}
