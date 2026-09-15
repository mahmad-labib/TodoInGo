package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *TodoHandler) Item(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		handleRequestError(w, ErrInvalidId)
		return
	}

	item, err := h.svc.FindByID(r.Context(), uint(val))
	if err != nil {
		handleServiceError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
