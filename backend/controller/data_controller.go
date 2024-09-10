package controller

import (
	"encoding/json"
	"g-now/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// のちにcontroller分割する
func GetAllData(c echo.Context) error {
	return c.JSON(http.StatusOK, "{status: 'ok'}")
}

func GetAllCategories(c echo.Context) error {
	// 仮実装：固定値
	categoryStr := `
	[
	  {"id": 1, "category_name": "政治"},
		{"id": 2, "category_name": "ビジネス"}
	]
	`

	var categories []model.Category
	err := json.Unmarshal([]byte(categoryStr), &categories)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse JSON"})
	}

	return c.JSON(http.StatusOK, categories)
}
