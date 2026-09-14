package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/service/validation"
	"gorm.io/gorm"
)

var ErrInvalidId = errors.New("id should be an integer")

type TodoService interface {
	Create(ctx context.Context, item *domain.Todo) (uint, error)
	List(ctx context.Context, sort string) ([]domain.TodoList, error)
	UpdateStatus(ctx context.Context, item *domain.Todo) (*domain.Todo, error)
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*domain.Todo, error)
	Update(ctx context.Context, item *domain.Todo) error
}

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

type TodoHandler struct {
	svc TodoService
}

func writeError(w http.ResponseWriter, status int, message string, details []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{
		Code:    status,
		Message: message,
		Details: details,
	})
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
		writeError(w, http.StatusUnprocessableEntity, "Validation Error", []string{err.Error()})
	case errors.Is(err, gorm.ErrRecordNotFound):
		writeError(w, http.StatusNotFound, "Not Found", []string{err.Error()})

	default:
		writeError(w, http.StatusInternalServerError, "internal server error", []string{})
	}
}

func handleRequestError(w http.ResponseWriter, err error) {
	writeError(w, http.StatusBadRequest, "Bad Request", []string{err.Error()})
}

func NewTodoRouter(svc TodoService) chi.Router {
	h := &TodoHandler{svc: svc}

	r := chi.NewRouter()
	r.Post("/v1/todo", h.Create)
	r.Get("/v1/todos", h.List)
	r.Get("/v1/todo/{id}", h.Item)
	r.Put("/v1/todo/{id}", h.Update)
	r.Patch("/v1/todo/{id}/status", h.UpdateStatus)
	r.Delete("/v1/todo/{id}", h.Delete)

	return r
}
