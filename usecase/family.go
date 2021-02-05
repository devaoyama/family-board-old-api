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
	user, err = fu.ur.AppendFamily(user, family)
	if err != nil {
		return o, err
	}
	o.Family = family

	return o, nil
}

func generateInvitationCode(length int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
