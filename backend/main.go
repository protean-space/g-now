package main

import (
	"g-now/controller"
	"g-now/db"
	"g-now/repository"
	"g-now/router"
	"g-now/usecase"
	"log/slog"
)

func main() {
	var err error
	dbConn, err := db.NewDB()
	if err != nil {
		slog.Error("fail connect db.")
		panic(err)
	}
	defer db.CloseDB(dbConn, err)

	// category
	categoryRepository := repository.NewCategoryRepository(dbConn)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUsecase)

	// article
	articleRepository := repository.NewArticleRepository(dbConn)
	articleUsecase := usecase.NewArticleUsecase(articleRepository)
	articleController := controller.NewArticleController(articleUsecase)

	// articles by category
	articlesByCategoryRepository := repository.NewArticlesByCategoryRepository(dbConn)
	articlesByCategoryUsecase := usecase.NewArticlesByCategoryUsecase(articlesByCategoryRepository)
	articlesByCategoryController := controller.NewArticlesByCategoryController(articlesByCategoryUsecase)

	// articles by tag
	articlesByTagRepository := repository.NewArticlesByTagRepository(dbConn)
	articlesByTagUsecase := usecase.NewArticlesByTagUsecase(articlesByTagRepository)
	articlesByTagController := controller.NewArticlesByTagController(articlesByTagUsecase)

	e := router.NewRouter(categoryController, articleController, articlesByCategoryController, articlesByTagController)
	e.Logger.Fatal(e.Start(":8080"))
}
