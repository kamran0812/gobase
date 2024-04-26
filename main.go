package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/kamran0812/gobase/handlers"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env")
	}
	router := chi.NewMux()
	router.Handle("/*", public())

	router.Get("/", handlers.Make(handlers.HandleHome))

	lisentAddress := os.Getenv("LISTEN_ADDR")
	slog.Info("Http server started", "listenAddress", lisentAddress)
	http.ListenAndServe(lisentAddress, router)
}
