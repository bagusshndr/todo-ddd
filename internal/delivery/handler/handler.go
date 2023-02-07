package handler

import (
	"test-golang/internal/shared"
	"test-golang/internal/usecase"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecaseTodo usecase.TodoUsecase
}

func (h *handler) HandlerGetTodo(c echo.Context) error {
	_, err := h.usecaseTodo.GetTodo()
	if err != nil {
		return shared.NewReponse("Gagal", 400, "Failed", nil, nil).JSON(c)
	}

	return shared.NewReponse("Sukses", 200, "Sukses", nil, nil).JSON(c)
}
