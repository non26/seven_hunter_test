package handler

import (
	"backend/third_find_all_meat/internal/core/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

type meathandler struct {
	meatUsecase port.IMeatUsecase
}

func NewMeatHandler(meatUsecase port.IMeatUsecase) *meathandler {
	return &meathandler{meatUsecase: meatUsecase}
}

func (h *meathandler) GetAllMeat(c echo.Context) error {
	meats, err := h.meatUsecase.GetAll(c.Request().Context())
	if err != nil {
		response := map[string]string{}
		response["error"] = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, meats)
}
