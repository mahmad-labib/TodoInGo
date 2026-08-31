package domain

import "time"

type Todo struct {
	ID        uint
	Title     string
	Status    string
	Priority  string
	CreatedAt time.Time
}
