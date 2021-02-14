package output

import "family-board-api/domain/model"

type GetTodo struct {
	Todos []*model.Todo
}

type CreateTodo struct {
	Result int
}

type ChangeTodoStatus struct {
	Result int
}

type DeleteTodo struct {
	Result int
}
