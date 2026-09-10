package domain

import "time"

type TodoList struct {
	ID        uint
	Title     string
	Status    string
	Priority  string
	CreatedAt time.Time
}
