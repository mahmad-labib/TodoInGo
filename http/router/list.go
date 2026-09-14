package router

import (
	"encoding/json"
	"net/http"
)

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")

	todoList, err := h.svc.List(r.Context(), sort)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoList)
}
