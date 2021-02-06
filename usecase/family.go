package usecase

import (
	"family-board-api/domain/model"
	"family-board-api/domain/repository"
	"family-board-api/usecase/input"
	"family-board-api/usecase/output"
	"math/rand"
)

type familyUsecase struct {
	ur repository.UserRepository
	fr repository.FamilyRepository
}

func NewFamilyUsecase(ur repository.UserRepository, fr repository.FamilyRepository) *familyUsecase {
	return &familyUsecase{
		ur: ur,
		fr: fr,
	}
}

func (fu *familyUsecase) CreateFamily(i *input.CreateFamily) (*output.CreateFamily, error) {
	o := &output.CreateFamily{}
	// Familyを作成
	family := model.NewFamily(i.Name, generateInvitationCode(8))
	family, err := fu.fr.Store(family)
	if err != nil {
		return o, err
	}
	user, err := fu.ur.FindById(i.UserId)
	if err != nil {
		return o, err
	}
	_, err = fu.fr.AppendUser(family, user)
	if err != nil {
		return o, err
	}
	o.Family = family

	return o, nil
}

func (fu *familyUsecase) JoinFamily(i *input.JoinFamily) (*output.JoinFamily, error) {
	o := &output.JoinFamily{}
	family, err := fu.fr.FindByInvitationCode(i.InvitationCode)
	if err != nil {
		return o, err
	}
	user, err := fu.ur.FindById(i.UserId)
	if err != nil {
		return o, err
	}
	_, err = fu.fr.AppendUser(family, user)
	o.Family = family

	return o, err
}

func (fu *familyUsecase) LeaveFamily(i *input.LeaveFamily) (*output.LeaveFamily, error) {
	o := &output.LeaveFamily{}
	family, err := fu.fr.FindById(i.FamilyId)
	if err != nil {
		return o, err
	}
	user, err := fu.ur.FindById(i.UserId)
	_, err = fu.fr.ClearUser(family, user)
	o.Result = 1

	return o, err
}

func generateInvitationCode(length int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
