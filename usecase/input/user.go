package input

type LoginWithLine struct {
	LiffIdToken string
}

type GetLoginUser struct {
	UserId int
}

type ChangeUserStatus struct {
	Status string
	UserId int
}
