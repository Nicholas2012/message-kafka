package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type StatisticsResponse struct {
	Total     int `json:"total"`
	Processed int `json:"processed"`
}

func (a *API) Statistics(w http.ResponseWriter, r *http.Request) {

	stat, err := a.service.Statistics(r.Context())
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	if err := json.NewEncoder(w).Encode(StatisticsResponse(*stat)); err != nil {
		slog.Error("Failed to write response", "error", err)
		return
	}
}
