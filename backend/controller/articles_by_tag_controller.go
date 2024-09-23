package controller

import (
	"g-now/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IArticlesByTagController interface {
	GetArticlesByTag(c echo.Context) error
}

type articlesByTagController struct {
	atu usecase.IArticlesByTagUsecase
}

func NewArticlesByTagController(atu usecase.IArticlesByTagUsecase) IArticlesByTagController {
	return &articlesByTagController{atu}
}

func (atc *articlesByTagController) GetArticlesByTag(c echo.Context) error {
	// パスパラメータの取得
	tagNameStr := c.Param("tagName")

	// usecase呼び出し: articlesの取得
	articles, err := atc.atu.GetArticlesByTag(tagNameStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}
