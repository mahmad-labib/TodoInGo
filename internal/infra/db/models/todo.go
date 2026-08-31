package models

import (
	"time"
)

type TodoModel struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"type:varchar(255);not null"`
	Status    string `gorm:"type:varchar(15);not null"`
	Priority  string `gorm:"type:varchar(15);not null"`
	CreatedAt time.Time
}

func (TodoModel) TableName() string {
	return "todo"
}
