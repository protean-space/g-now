package main

import (
	"g-now/db"
	"g-now/model"
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
	if err != nil {
		slog.Error("fail close db connect.")
		panic(err)
	}

	dbConn.AutoMigrate(&model.Article{}, &model.Category{}, &model.ArticleCategoryMap{})
	slog.Info("Successfully Migrated.")
}
