package postgresql

import (
	"fmt"
)

type RentalRepository struct {
	*repository
}

func NewRentalRepository() *RentalRepository {
	repo := newRepository()

	return &RentalRepository{
		repo,
	}
}

func (r *RentalRepository) FindById(id string) (FindResult, error) {
	result := FindResult{}
	err := r.db.Model(&Rentals{}).Select("*").
		Joins("join users on users.id = rentals.user_id").Where("rentals.id = ?", id).Scan(&result).Error

	if err != nil {
		return result, fmt.Errorf("error while fetching from database: %w", err)
	}

	return result, nil
}
