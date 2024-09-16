package controller

import (
	"g-now/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IArticleController interface {
	GetAllArticles(c echo.Context) error
}

type articleController struct {
	au usecase.IArticleUsecase
}

func NewArticleController(au usecase.IArticleUsecase) IArticleController {
	return &articleController{au}
}

func (ac *articleController) GetAllArticles(c echo.Context) error {
	articles, err := ac.au.GetAllArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}
