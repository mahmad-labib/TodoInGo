package router

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		handleRequestError(w, ErrInvalidId)
		return
	}

	err = h.svc.Delete(r.Context(), uint(val))
	if err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
