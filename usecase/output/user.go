package output

import "family-board-api/domain/model"

type LoginWithLine struct {
	JwtToken string
}

type GetLoginUser struct {
	User *model.User
}

type ChangeUserStatus struct {
	User *model.User
}
