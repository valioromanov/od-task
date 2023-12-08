package postgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db gorm.DB
}

func NewRepository() (*Repository, error) {
	repository := Repository{}
	dsn := "host=localhost user=root password=root dbname=testingwithrentals port=5434 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return &repository, err
	}

	repository.db = *db
	return &repository, nil
}
