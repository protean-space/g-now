package controller

import (
	"g-now/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ICategoryController interface {
	GetAllCategories(c echo.Context) error
}

type categoryController struct {
	cu usecase.ICategoryUsecase
}

func NewCategoryController(cu usecase.ICategoryUsecase) ICategoryController {
	return &categoryController{cu}
}

func (cc *categoryController) GetAllCategories(c echo.Context) error {
	categoryRes, err := cc.cu.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, categoryRes)
}
