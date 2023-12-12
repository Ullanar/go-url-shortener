package routes

import (
	"math/rand"
	"net/http"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/repository"
)

func CreateAlias(w http.ResponseWriter, r *http.Request, repository repository.Repository) {
	cfg := config.MustLoad()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	alias := RandomizeAlias(8)

	createErr := repository.CreateAlias(alias, r.Form.Get("dest"))
	if createErr != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Something went wrong"))
		return
	}

	_, _ = w.Write([]byte(makeResponseString(cfg, alias)))
}

func makeResponseString(cfg *config.Config, alias string) string {

	response := ""
	if cfg.Server.UseHttps {
		response += "https://"
	} else {
		response += "http://"
	}
	response += cfg.Server.Host
	if cfg.Server.IncludePort {
		response += cfg.Server.Port
	}
	response += "/" + alias

	return response
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomizeAlias(length int) string {
	return StringWithCharset(length, charset)
}
