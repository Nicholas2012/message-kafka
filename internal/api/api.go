package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"testkafka/internal/models"
)

type Service interface {
	CreateMessage(ctx context.Context, message string) error
	Statistics(ctx context.Context) (*models.Statistics, error)
}

type API struct {
	service Service
}

func New(s Service) *API {
	return &API{
		service: s,
	}
}

func (a *API) AddRoutes(s *http.ServeMux) {
	s.HandleFunc("POST /messages", a.CreateMessage)
	s.HandleFunc("GET /statistics", a.Statistics)
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func writeErr(w http.ResponseWriter, r *http.Request, status int, err error) {
	slog.Error("Request failed", "status", status, "error", err.Error(), "url", r.URL.Path)

	resp := ErrorResponse{
		Error: err.Error(),
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("Failed to write response", "error", err)
		return
	}
}

func badRequest(w http.ResponseWriter, r *http.Request, err error) {
	writeErr(w, r, http.StatusBadRequest, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	writeErr(w, r, http.StatusInternalServerError, err)
}
