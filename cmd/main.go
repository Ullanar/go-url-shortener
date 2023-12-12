package main

import (
	"log/slog"
	"net/http"
	"os"
	"url-shortener/api"
	"url-shortener/internal/app"
	"url-shortener/internal/logger"
)

func main() {
	provider := app.New()
	log := logger.New(provider.Config().Env)
	log.Info("Config and Logger was initialized", slog.String("env", provider.Config().Env))

	router := api.New(provider.Repository())

	err := http.ListenAndServe(provider.Config().Server.Port, router)
	if err != nil {
		log.Error("Error on server start:", err)
		os.Exit(1)
	}
}
