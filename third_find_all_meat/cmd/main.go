package main

import (
	"backend/third_find_all_meat/internal/core/usecases"
	"backend/third_find_all_meat/internal/handler"
	"backend/third_find_all_meat/internal/repositories"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	meatRepository := repositories.NewMeatRepository()
	meatUsecase := usecases.NewMeatUsecase(meatRepository)
	meatHandler := handler.NewMeatHandler(meatUsecase)
	e.GET("/meats", meatHandler.GetAllMeat)
	e.Start(":8080")
}
