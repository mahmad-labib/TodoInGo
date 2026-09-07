package domain

import "github.com/mahmad-labib/TodoInGo/internal/infra/db/model"

type Sort struct {
	Sort string
}

type List struct {
	Data []model.Todo
}
