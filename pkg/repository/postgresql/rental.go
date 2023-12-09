package postgresql

import (
	"fmt"
	"od-task/pkg/helper"
	"strconv"

	"gorm.io/gorm"
)

const WITHIN_MILES int = 100

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
	tx := r.db.Table("rentals").Select("*").
		Joins("join users on users.id = rentals.user_id").Where("rentals.id = ?", id).Find(&result)

	if tx.Error != nil {
		return FindResult{}, fmt.Errorf("error while fetching from database: %w", tx.Error)
	}

	if tx.RowsAffected == 0 {
		return FindResult{}, fmt.Errorf("no rows affected")
	}
	fmt.Println("id: ", result.ID)
	return result, nil
}

func (r *RentalRepository) FindByFilters(filters map[string][]string) ([]FindResult, error) {
	results := []FindResult{}

	query, err := buildDynamicQueryByFilters(r, filters)
	if err != nil {
		return nil, fmt.Errorf("error while building query: %w", err)
	}

	query.Find(&results)

	return results, nil
}

func buildDynamicQueryByFilters(repo *RentalRepository, filters map[string][]string) (*gorm.DB, error) {
	query := repo.db.Table("rentals").Select("*").
		Joins("join users on  rentals.user_id = users.id")

	for key, val := range filters {

		switch key {
		case "limit":
			limit, err := strconv.Atoi(val[0])
			if err != nil {
				return nil, fmt.Errorf("not proper value for limit")
			}
			query.Limit(limit)
		case "offset":
			offset, err := strconv.Atoi(val[0])
			if err != nil {
				return nil, fmt.Errorf("not proper value for offset")
			}
			query.Offset(offset)
		case "ids":
			ids, err := helper.StringToIntArray(val[0])
			if err != nil {
				return nil, err
			}
			query.Where("rentals.id in ?", ids)
		case "price_max":
			query.Where("price_per_day <= ?", val)
		case "price_min":
			query.Where("price_per_day >= ?", val)
		case "sort":
			query.Order(fmt.Sprintf("%s asc", val[0]))
		case "near":
			points, err := helper.StringToFloat64Array(val[0])
			if err != nil {
				return nil, err
			}
			if len(points) != 2 {
				return nil, fmt.Errorf("not proper count values for near: recieved %d wants %d ", len(points), 2)
			}
			query.Where("ST_DWithin(ST_MakePoint(rentals.lat,rentals.lng), ST_MakePoint(?,?)::geography, ? * 1609.34)", points[0], points[1], WITHIN_MILES)
		default:
			query.Where(fmt.Sprintf("%s = ?", key), val)
		}
	}
	return query, nil
}
