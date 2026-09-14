package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

func (h *TodoHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var body UpdateStatus
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleRequestError(w, err)
		return
	}

	id := chi.URLParam(r, "id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		handleRequestError(w, ErrInvalidId)
		return
	}

	item := domain.Todo{
		ID:     uint(val),
		Status: body.Status,
	}

	todo, err := h.svc.UpdateStatus(r.Context(), &item)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}
