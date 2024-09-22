package controller

import (
	"g-now/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IArticlesByCategoryController interface {
	GetArticlesByCategory(c echo.Context) error
}

type articlesByCategoryController struct {
	acu usecase.IArticlesByCategoryUsecase
}

func NewArticlesByCategoryController(acu usecase.IArticlesByCategoryUsecase) IArticlesByCategoryController {
	return &articlesByCategoryController{acu}
}

func (acr *articlesByCategoryController) GetArticlesByCategory(c echo.Context) error {
	// パスパラメータの取得
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	// usecase呼び出し: articlesの取得
	articles, err := acr.acu.GetArticlesByCategory(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}
