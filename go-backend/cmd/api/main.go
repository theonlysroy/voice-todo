package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/theonlysroy/voice-todo/go-backend/internal/middleware"
)

type ApiResponse struct {
	Messagge string `json:"message"`
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// db connection
	// pool, err := pgxpool.New(ctx, os.Getenv(DB_URL))

	// todoSvc := todo.NewService()
	r := chi.NewRouter()
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&ApiResponse{
			Messagge: "Api ok...",
		})
	})

	r.Use(middleware.ReqLogger)
	r.Route("/api", func(r chi.Router) {
	})

	srv := &http.Server{
		Addr:         ":4040",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 35 * time.Second,
	}

	slog.Info("listening", "addr", srv.Addr)
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
