package routes

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/database"
)

func CreateAlias(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	cfg := config.MustLoad()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	newAlias := database.Link{Alias: RandomizeAlias(8), Dest: r.Form.Get("dest")}
	result := db.Create(&newAlias)

	if result.Error != nil {
		w.WriteHeader(500)
		fmt.Println(result.Error)
		_, _ = w.Write([]byte("Something went wrong"))
		return
	}

	_, _ = w.Write([]byte(makeResponseString(cfg, newAlias.Alias)))
	return
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
