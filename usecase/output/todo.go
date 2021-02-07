package output

import "family-board-api/domain/model"

type GetTodo struct {
	Todos []*model.Todo
}

type CreateTodo struct {
	Todo *model.Todo
}

type ChangeTodoStatus struct {
	Todo *model.Todo
}

type DeleteTodo struct {
	Todo *model.Todo
}
