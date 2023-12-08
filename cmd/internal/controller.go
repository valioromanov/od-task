package internal

import "od-task/pkg/repository/postgresql"

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

func (c *Controller) GetVehicleByID(id string) (VehicleInfo, Location, User) {
	c.repo.FindById(id)
	return VehicleInfo{}, Location{}, User{}
}
