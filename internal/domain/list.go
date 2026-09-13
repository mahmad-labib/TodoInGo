package domain

import "time"

type TodoList struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
}
