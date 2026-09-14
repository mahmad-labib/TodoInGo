package router

import (
	"encoding/json"
	"net/http"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleRequestError(w, err)
		return
	}

	todo := domain.Todo{
		Title:    body.Title,
		Status:   body.Status,
		Priority: body.Priority,
	}

	id, err := h.svc.Create(r.Context(), &todo)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{"id": id})
}
