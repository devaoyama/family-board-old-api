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

type loginWithLiffRequest struct {
	IdToken string `json:"id_token"`
}

type changeStatusRequest struct {
	Status string `json:"status"`
}

func (uh *UserHandler) LoginWithLiff(ctx echo.Context) error {
	ur := uh.repository.Ur
	uu := usecase.NewUserUsecase(ur)

	var request loginWithLiffRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	i := &input.LoginWithLine{LiffIdToken: request.IdToken}
	o, err := uu.LoginWithLine(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, o.JwtToken)
}

func (uh *UserHandler) Me(ctx echo.Context) error {
	ur := uh.repository.Ur
	uu := usecase.NewUserUsecase(ur)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	i := &input.GetLoginUser{
		UserId: userId,
	}
	o, err := uu.GetLoginUser(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}

func (uh *UserHandler) ChangeStatus(ctx echo.Context) error {
	ur := uh.repository.Ur
	uu := usecase.NewUserUsecase(ur)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	var request changeStatusRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	i := &input.ChangeUserStatus{
		Status: request.Status,
		UserId: userId,
	}
	o, err := uu.ChangeUserStatus(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}
