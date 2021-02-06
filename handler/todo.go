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

type TodoHandler struct {
	repository *registry.Repository
}

func NewTodoHandler(repository *registry.Repository) *TodoHandler {
	return &TodoHandler{repository: repository}
}

func (th *TodoHandler) Create(ctx echo.Context) error {
	ur := th.repository.Ur
	fr := th.repository.Fr
	tr := th.repository.Tr
	tu := usecase.NewTodoUsecase(ur, fr, tr)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	i := &input.CreateTodo{
		Title:       ctx.FormValue("title"),
		Description: ctx.FormValue("description"),
		UserId:      userId,
	}
	o, err := tu.CreateTodo(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}

func (th *TodoHandler) ChangeStatus(ctx echo.Context) error {
	ur := th.repository.Ur
	fr := th.repository.Fr
	tr := th.repository.Tr
	tu := usecase.NewTodoUsecase(ur, fr, tr)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	todoId, _ := strconv.Atoi(ctx.Param("id"))
	i := &input.ChangeTodoStatus{
		UserId: userId,
		TodoId: todoId,
	}
	o, err := tu.ChangeTodoStatus(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}

func (th *TodoHandler) Delete(ctx echo.Context) error {
	ur := th.repository.Ur
	fr := th.repository.Fr
	tr := th.repository.Tr
	tu := usecase.NewTodoUsecase(ur, fr, tr)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	userId := claims.Id

	todoId, _ := strconv.Atoi(ctx.Param("id"))
	i := &input.DeleteTodo{
		UserId: userId,
		TodoId: todoId,
	}
	o, err := tu.DeleteTodo(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, o)
}
