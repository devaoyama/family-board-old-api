package input

type GetTodo struct {
	UserId int
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
