package db

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"go.uber.org/multierr"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	var db *gorm.DB
	var url string
	var err error
	if os.Getenv("GO_ENV") == "dev" {
		slog.Info("GO_ENV: dev")
		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	} else {
		slog.Info("GO_ENV: production")

		dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PW"))
		db, err = gorm.Open(postgres.New(postgres.Config{DriverName: "cloudsqlpostgres", DSN: dsn}))
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
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
