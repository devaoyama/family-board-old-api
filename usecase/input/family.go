package input

type CreateFamily struct {
	Name   string
	UserId int
}

type JoinFamily struct {
	InvitationCode string
	UserId         int
}