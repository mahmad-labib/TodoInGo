package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/service"
)

type TodoBody struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
}

type TodoHandler struct {
	svc *service.Todo
	ctx context.Context
}

func NewTodoRouter(ctx context.Context, svc *service.Todo) chi.Router {
	h := &TodoHandler{svc: svc, ctx: ctx}

	r := chi.NewRouter()
	r.Post("/", h.Create)

	return r
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body TodoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item := domain.Todo{
		Title:    body.Title,
		Status:   body.Status,
		Priority: body.Priority,
	}

	id, err := h.svc.Create(h.ctx, &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if id > 0 {
		w.WriteHeader(http.StatusCreated)
		w.Write(fmt.Appendf([]byte("todo created with id: "), "%d", id))
	}
}
