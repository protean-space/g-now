package db

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"go.uber.org/multierr"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	if os.Getenv("GO_ENV") == "dev" {
		if err := godotenv.Load(); err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	slog.Info("DB Connected.")
	return db, nil
}

func CloseDB(db *gorm.DB, err error) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		err = multierr.Append(err, errors.New("fail close db connect."))
	}
}
