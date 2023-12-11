package postgresql

import (
	"fmt"
	"od-task/cmd/env"
	"od-task/pkg/app"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	db gorm.DB
}

func newRepository(config env.AppConfig) *repository {
	repository := repository{}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Error while connection to the database: ", err)
		app.Crash(err)
	}
	logrus.Info("Successfull connection to the database")
	repository.db = *db
	return &repository
}
