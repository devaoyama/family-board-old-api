package input

type CreateTodo struct {
	Title       string
	Description string
	FamilyId    int
	UserId      int
}

type ChangeTodoStatus struct {
	TodoId int
	UserId int
}
