package handler

import (
	"family-board-api/auth"
	"family-board-api/registry"
	"family-board-api/usecase"
	"family-board-api/usecase/input"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type FamilyHandler struct {
	repository *registry.Repository
}

func NewFamilyHandler(repository *registry.Repository) *FamilyHandler {
	return &FamilyHandler{repository: repository}
}

func (fh *FamilyHandler) Get(ctx echo.Context) error {
	ur := fh.repository.Ur
	fr := fh.repository.Fr
	fu := usecase.NewFamilyUsecase(ur, fr)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	i := &input.GetFamily{
		UserId: userId,
	}
	o, err := fu.GetFamily(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}

func (fh *FamilyHandler) Create(ctx echo.Context) error {
	ur := fh.repository.Ur
	fr := fh.repository.Fr
	fu := usecase.NewFamilyUsecase(ur, fr)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	i := &input.CreateFamily{
		Name: ctx.FormValue("name"),
		UserId: userId,
	}
	o, err := fu.CreateFamily(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}

func (fh *FamilyHandler) Join(ctx echo.Context) error {
	ur := fh.repository.Ur
	fr := fh.repository.Fr
	fu := usecase.NewFamilyUsecase(ur, fr)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	i := &input.JoinFamily{
		InvitationCode: ctx.FormValue("invitation_code"),
		UserId: userId,
	}
	o, err := fu.JoinFamily(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}

func (fh *FamilyHandler) Leave(ctx echo.Context) error {
	ur := fh.repository.Ur
	fr := fh.repository.Fr
	fu := usecase.NewFamilyUsecase(ur, fr)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	familyId, _ := strconv.Atoi(ctx.Param("id"))
	i := &input.LeaveFamily{
		UserId: userId,
		FamilyId: familyId,
	}
	o, err := fu.LeaveFamily(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}
