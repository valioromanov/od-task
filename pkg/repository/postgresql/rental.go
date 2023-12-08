package postgresql

type RentalRepository struct {
	*repository
}

func NewRentalRepository() *RentalRepository {
	repo := NewRepository()

	return &RentalRepository{
		repo,
	}
}

func (r *RentalRepository) FindById(id string) (FindResult, error) {
	result := FindResult{}
	err := r.db.Model(&Rentals{}).Select("rentals.*, users.* ").
		Joins("join users on users.id = rentals.user_id").Where("rentals.id = ?", id).Scan(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
