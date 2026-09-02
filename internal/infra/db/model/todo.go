package model

import (
	"time"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Status    string    `gorm:"type:varchar(15);not null"`
	Priority  string    `gorm:"type:varchar(15);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP"`
}

func (Todo) TableName() string {
	return "todo"
}
