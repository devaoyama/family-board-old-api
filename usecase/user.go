package usecase

import (
	"errors"
	"family-board-api/auth"
	"family-board-api/domain/repository"
	"family-board-api/line"
	"family-board-api/usecase/input"
	"family-board-api/usecase/output"

	"gorm.io/gorm"
)

type UserUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) *UserUsecase {
	return &UserUsecase{ur: ur}
}

func (uu *UserUsecase) LoginWithLine(i *input.LoginWithLine) (*output.LoginWithLine, error) {
	// LIFFのTokenの有効性を確認してUIDを取得
	liff := line.VerifiedIdToken(i.LiffIdToken)

	// UIDからUserを作成or更新
	user, err := uu.ur.FindByLineUid(liff.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 作成
			user.LineUid = liff.Uid
			user.Name = liff.Name
			user.PictureUrl = liff.Picture
			user, err = uu.ur.Store(user)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// 更新
		user.Name = liff.Name
		user.PictureUrl = liff.Picture
		user, err = uu.ur.Update(user)
		if err != nil {
			return nil, err
		}
	}

	// jwtトークンを発行して返却
	jwtToken, err := auth.GenerateJwtToken(int(user.ID), user.Name, user.PictureUrl)
	o := &output.LoginWithLine{
		JwtToken: jwtToken,
	}

	return o, err
}

func (uu *UserUsecase) ChangeUserStatus(i *input.ChangeUserStatus) (*output.ChangeUserStatus, error) {
	o := &output.ChangeUserStatus{}
	user, err := uu.ur.FindById(i.UserId)
	if err != nil {
		return o, err
	}
	user.Status = i.Status
	user, err = uu.ur.Update(user)
	o.User = user
	return o, err
}
