package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/theonlysroy/voice-todo/go-backend/internal/httpx"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

type createRequest struct {
	Title string `json:"title" doc:"Todo title is required"`
}

type updateRequest struct {
	Title       string `json:"title" doc:"Todo title is required"`
	IsCompleted string `json:"isCompleted" doc:"Is todo complete?"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	// validate
	t, err := h.svc.Create(r.Context(), req.Title)
	if err != nil {
		return err
	}
	httpx.JSON(w, http.StatusCreated, t)
	return nil
}

func idParam(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := idParam(r)
	if err != nil {
		return httpx.BadRequest("Invalid id", id)
	}
	t, err := h.svc.Get(r.Context(), id)
	if err != nil {
		return httpx.NotFound("todo not found")
	}
	httpx.JSON(w, http.StatusOK, t)
	return nil
}
