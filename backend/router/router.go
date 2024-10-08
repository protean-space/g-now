package router

import (
	"g-now/controller"
	"g-now/cron"
	"g-now/db/seed"
	"g-now/migrate"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	cc controller.ICategoryController,
	ac controller.IArticleController,
	acc controller.IArticlesByCategoryController,
	atc controller.IArticlesByTagController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
	}))

	e.GET("/categories", cc.GetAllCategories)
	e.GET("/articles", ac.GetAllArticles)
	e.GET("/categories/:id/articles", acc.GetArticlesByCategory)
	e.GET("/tags/:tagName/articles", atc.GetArticlesByTag)
	e.GET("/migrate_and_seed", migrateAndSeed)
	e.GET("/fetch_news", fetchNews)

	return e
}

func migrateAndSeed(c echo.Context) error {
	migrate.Run()
	slog.Info("finish migrate.")
	seed.Run()
	slog.Info("finish seed.")
	return nil
}

func fetchNews(c echo.Context) error {
	cron.FetchNews()
	slog.Info("finish fetch Google News.")
	return nil
}
