package output

import "family-board-api/domain/model"

type LoginWithLine struct {
	JwtToken string
}

type ChangeUserStatus struct {
	User *model.User
}
