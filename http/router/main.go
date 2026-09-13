package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/service/validation"
)

type CreateTodoRequest struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
}

type TodoService interface {
	Create(ctx context.Context, item *domain.Todo) (uint, error)
	List(ctx context.Context, sort string) ([]domain.TodoList, error)
	UpdateStatus(ctx context.Context, item *domain.Todo) (*domain.Todo, error)
	Delete(ctx context.Context, id uint) error
}

type TodoHandler struct {
	svc TodoService
}

func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, validation.ErrInvalidTitle),
		errors.Is(err, validation.ErrTitleLength),
		errors.Is(err, validation.ErrInvalidPriority),
		errors.Is(err, validation.ErrPriorityLength),
		errors.Is(err, validation.ErrInvalidStatus),
		errors.Is(err, validation.ErrStatusLength),
		errors.Is(err, validation.ErrInvalidSort):
		http.Error(w, err.Error(), http.StatusBadRequest)

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func NewTodoRouter(svc TodoService) chi.Router {
	h := &TodoHandler{svc: svc}

	r := chi.NewRouter()
	r.Post("/todo", h.Create)
	r.Get("/todo", h.List)

	return r
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleServiceError(w, err)
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
