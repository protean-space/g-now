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

	categoryRepository := repository.NewCategoryRepository(dbConn)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUsecase)

	e := router.NewRouter(categoryController)
	e.Logger.Fatal(e.Start(":8080"))
}
