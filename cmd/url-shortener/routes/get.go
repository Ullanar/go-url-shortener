package routes

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"url-shortener/internal/config"
	"url-shortener/internal/database"
)

type GetDestResponse struct {
	dest string
}

func GetDestAndRedirect(w http.ResponseWriter, r *http.Request) {
	cfg := config.MustLoad()
	db := database.New(cfg.Database)

	alias := chi.URLParam(r, "alias")

	res, err := db.Query(fmt.Sprintf(`SELECT dest FROM links WHERE alias = '%s'`, alias))
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Something went wrong"))
		return
	}

	var response GetDestResponse
	for res.Next() {
		err := res.Scan(&response.dest)
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("Something went wrong"))
			_ = db.Close()
			return
		}
	}
	_ = db.Close()

	if response.dest == "" {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("Not found"))
		return
	}

	http.Redirect(w, r, response.dest, http.StatusFound)
	return
}
