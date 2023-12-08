package internal

import (
	"od-task/pkg/repository/postgresql"
)

type RentalRepository interface {
	FindById(id string) (postgresql.FindResult, error)
}

type Controller struct {
	repo RentalRepository
}

func NewController(repository RentalRepository) *Controller {
	return &Controller{
		repository,
	}
}

func (c *Controller) GetVehicleByID(id string) (VehicleInfo, Location, User, error) {
	var vehicle = postgresql.FindResult{}
	vehicle, err := c.repo.FindById(id)

	if err != nil {
		return VehicleInfo{}, Location{}, User{}, err
	}
	vehicleInfo, location, user := findResultToControllerResponse(vehicle)
	return vehicleInfo, location, user, nil
}

func findResultToControllerResponse(vehicle postgresql.FindResult) (VehicleInfo, Location, User) {
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
		Latitude:   vehicle.Lat,
		Longtitude: vehicle.Lng,
	}

	var user = User{
		UserID:    vehicle.UserID,
		FirstName: vehicle.FirstName,
		LastName:  vehicle.LastName,
	}

	return vehicleInfo, location, user
}
