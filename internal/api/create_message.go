package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CreateMessageRequest struct {
	Message string `json:"message"`
}

func (a *API) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req CreateMessageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, r, err)
		return
	}

	if len(req.Message) == 0 {
		badRequest(w, r, errors.New("Message can't be empty!"))
		return
	}

	if err := a.service.CreateMessage(r.Context(), req.Message); err != nil {
		internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
