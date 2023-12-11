package postgresql

import (
	"fmt"
	"od-task/cmd/env"
	"od-task/pkg/helper"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const WITHIN_MILES int = 100

type RentalRepository struct {
	*repository
}

func NewRentalRepository(config env.AppConfig) *RentalRepository {
	repo := newRepository(config)

	return &RentalRepository{
		repo,
	}
}

func (r *RentalRepository) FindById(id string) (FindResult, error) {
	result := FindResult{}
	tx := r.db.Table("rentals").Select("*").
		Joins("join users on users.id = rentals.user_id").Where("rentals.id = ?", id).Find(&result)

	if tx.Error != nil {
		logrus.Error(fmt.Sprintf("error while fetching rental by id '%s' from database: %s ", id, tx.Error.Error()))
		return FindResult{}, fmt.Errorf("error while fetching from database: %w", tx.Error)
	}

	if tx.RowsAffected == 0 {
		logrus.Error("No data found for rentals with id: ", id)
		return FindResult{}, fmt.Errorf("no rows affected")
	}
	return result, nil
}

func (r *RentalRepository) FindByFilters(filters map[string][]string) ([]FindResult, error) {
	results := []FindResult{}

	query, err := buildDynamicQueryByFilters(r, filters)
	if err != nil {
		logrus.Error("Error while bulding an query for getting rentals by filters: ", err)
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
				logrus.Error("value for limit is not a number")
				return nil, fmt.Errorf("not proper value for limit")
			}
			query.Limit(limit)
		case "offset":
			offset, err := strconv.Atoi(val[0])
			if err != nil {
				logrus.Error("value for offset is not a number")
				return nil, fmt.Errorf("not proper value for offset")
			}
			query.Offset(offset)
		case "ids":
			ids, err := helper.StringToIntArray(val[0])
			if err != nil {
				logrus.Error("cannot parse string to int array: ", err)
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
				logrus.Error("cannot parse string to float64 array: ", err)
				return nil, err
			}
			if len(points) != 2 {
				logrus.Error(fmt.Sprintf("not proper count of values for 'near' received %d wants 2", len(points)))
				return nil, fmt.Errorf("not proper count values for near: recieved %d wants %d ", len(points), 2)
			}
			query.Where("ST_DWithin(ST_MakePoint(rentals.lat,rentals.lng), ST_MakePoint(?,?)::geography, ? * 1609.34)", points[0], points[1], WITHIN_MILES)
		default:
			query.Where(fmt.Sprintf("%s = ?", key), val)
		}
	}
	return query, nil
}
