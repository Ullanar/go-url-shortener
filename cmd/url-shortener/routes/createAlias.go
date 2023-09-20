package routes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/database"
)

type CreateAliasRequestBody struct {
	Dest string `json:"dest"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"

func CreateAlias(w http.ResponseWriter, r *http.Request) {
	cfg := config.MustLoad()
	db := database.New(cfg.Database)

	var body CreateAliasRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		_, _ = w.Write([]byte("Bad request body"))
		return
	}

	alias := String(8)
	_, err = db.Query(fmt.Sprintf("INSERT INTO links (dest, alias) VALUES ('%s', '%s')", body.Dest, alias))
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Something went wrong"))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("%s/%s", cfg.Server.Host, alias)))
	return
}

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
