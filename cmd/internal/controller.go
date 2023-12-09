package internal

import (
	"fmt"
	"od-task/pkg/repository/postgresql"
)

type RentalRepository interface {
	FindById(id string) (postgresql.FindResult, error)
	FindByFilters(filters map[string][]string) ([]postgresql.FindResult, error)
}

type Controller struct {
	repo RentalRepository
}

func NewController(repository RentalRepository) *Controller {
	return &Controller{
		repository,
	}
}

func (c *Controller) GetVehicleByID(id string) (GetRentalResponse, error) {
	var vehicle = postgresql.FindResult{}
	vehicle, err := c.repo.FindById(id)

	if err != nil {
		return GetRentalResponse{}, err
	}
	rental := findResultToControllerResponse(vehicle)
	return rental, nil
}

func (c *Controller) GetFilteredRentals(filters map[string][]string) ([]GetRentalResponse, error) {
	filteredRentals, err := c.repo.FindByFilters(filters)
	if err != nil {
		return nil, fmt.Errorf("error while fetching filtered rentals: %s", err.Error())
	}

	rentals := make([]GetRentalResponse, len(filteredRentals))

	for ind, val := range filteredRentals {
		rentals[ind] = findResultToControllerResponse(val)
	}

	return rentals, nil
}

func findResultToControllerResponse(vehicle postgresql.FindResult) GetRentalResponse {
	var vehicleInfo = VehicleInfo{
		VehicleID:       vehicle.ID,
		Name:            vehicle.Name,
		Description:     vehicle.Description,
		Type:            vehicle.Type,
		Make:            vehicle.VehicleMake,
		Model:           vehicle.VehicleModel,
		Year:            vehicle.VehicleYear,
		Length:          vehicle.VehicleLength,
		Sleeps:          vehicle.Sleeps,
		PrimaryImageUrl: vehicle.PrimaryImageUrl,
		Price: Price{
			Day: vehicle.PricePerDay,
		},
	}

	var location = Location{
		City:       vehicle.HomeCity,
		State:      vehicle.HomeState,
		Zip:        vehicle.HomeZip,
		Country:    vehicle.HomeCountry,
		Latitude:   vehicle.Latitude,
		Longtitude: vehicle.Longtitude,
	}

	var user = User{
		UserID:    vehicle.UserID,
		FirstName: vehicle.FirstName,
		LastName:  vehicle.LastName,
	}

	return GetRentalResponse{
		vehicleInfo,
		location,
		user,
	}
}
