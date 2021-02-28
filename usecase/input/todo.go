package input

import "time"

type GetTodo struct {
	UserId int
	Date   *time.Time
}

type CreateTodo struct {
	Title       string
	Description string
	UserId      int
}

type ChangeTodoStatus struct {
	TodoId int
	UserId int
}

type DeleteTodo struct {
	TodoId int
	UserId int
}
