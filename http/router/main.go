package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type TodoService interface {
	Create(ctx context.Context, item *domain.Todo) (uint, error)
	List(ctx context.Context, sort string) ([]domain.TodoList, error)
	UpdateStatus(ctx context.Context, item *domain.Todo) (*domain.Todo, error)
	Delete(ctx context.Context, id uint) error
}

type TodoHandler struct {
	svc TodoService
}

func NewTodoRouter(svc TodoService) chi.Router {
	h := &TodoHandler{svc: svc}

	r := chi.NewRouter()
	r.Post("/todo", h.Create)
	r.Get("/todo", h.List)

	return r
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.svc.Create(r.Context(), &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(fmt.Appendf([]byte("todo created with id: "), "%d", id))
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")

	todoList, err := h.svc.List(r.Context(), sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoList)
}
