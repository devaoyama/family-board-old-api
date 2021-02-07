package handler

import (
	"family-board-api/auth"
	"family-board-api/registry"
	"family-board-api/usecase"
	"family-board-api/usecase/input"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repository *registry.Repository
}

func NewUserHandler(repository *registry.Repository) *UserHandler {
	return &UserHandler{repository: repository}
}

func (uh *UserHandler) LoginWithLiff(ctx echo.Context) error {
	ur := uh.repository.Ur
	uu := usecase.NewUserUsecase(ur)

	i := &input.LoginWithLine{LiffIdToken: ctx.FormValue("id_token")}
	o, err := uu.LoginWithLine(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, o.JwtToken)
}

func (uh *UserHandler) ChangeStatus(ctx echo.Context) error {
	ur := uh.repository.Ur
	uu := usecase.NewUserUsecase(ur)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	i := &input.ChangeUserStatus{
		Status: ctx.FormValue("status"),
		UserId: userId,
	}
	o, err := uu.ChangeUserStatus(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}
