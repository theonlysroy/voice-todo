package todo

import (
	"github.com/go-chi/chi/v5"
	"github.com/theonlysroy/voice-todo/go-backend/internal/httpx"
)

func (h *Handler) Routes(r chi.Router) {
	r.Route("/todos", func(r chi.Router) {
		r.Post("/", httpx.Wrap(h.Create))
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", httpx.Wrap(h.Get))
		})
	})
}
