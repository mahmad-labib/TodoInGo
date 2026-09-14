package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		handleRequestError(w, ErrInvalidId)
		return
	}

	var body CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleRequestError(w, err)
		return
	}

	todo := domain.Todo{
		ID:       uint(val),
		Title:    body.Title,
		Status:   body.Status,
		Priority: body.Priority,
	}

	err = h.svc.Update(r.Context(), &todo)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
