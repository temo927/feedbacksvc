
package httpserver

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

    "github.com/temo927/feedbacksvc/internal/config"
    "github.com/temo927/feedbacksvc/internal/http/handlers"
    "github.com/temo927/feedbacksvc/internal/pubsub"
    "github.com/temo927/feedbacksvc/internal/store"
)

func NewRouter(st store.Store, pub pubsub.Publisher, cfg config.Config) *chi.Mux {
    r := chi.NewRouter()

    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    h := handlers.New(st, pub, cfg)

    r.Post("/feedback", h.CreateFeedback)
    r.Get("/health", h.Health)

    return r
}
