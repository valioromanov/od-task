package postgresql

import (
	"od-task/pkg/app"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	db gorm.DB
}

func newRepository() *repository {
	repository := repository{}
	dsn := "host=localhost user=root password=root dbname=testingwithrentals port=5434 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Error while connection to the database")
		app.Crash(err)
	}
	logrus.Info("Successfull connection to the database")
	repository.db = *db
	return &repository
}
