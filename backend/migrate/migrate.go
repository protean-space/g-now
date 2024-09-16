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

	dbConn.AutoMigrate(&model.Article{}, &model.Category{}, &model.ArticleCategoryMap{}, &model.Tag{}, &model.ArticleTagMap{})
	slog.Info("Successfully Migrated.")
}
