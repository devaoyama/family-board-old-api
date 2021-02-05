package output

import "family-board-api/domain/model"

type CreateTodo struct {
	Todo *model.Todo
}

type ChangeStatus struct {
	Todo *model.Todo
}
