package output

import "family-board-api/domain/model"

type GetFamily struct {
	Family *model.Family
}

type CreateFamily struct {
	Family *model.Family
}

type JoinFamily struct {
	Family *model.Family
}

type LeaveFamily struct {
	Result int
}
