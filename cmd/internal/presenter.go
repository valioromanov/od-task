package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RentalController interface {
	GetVehicleByID(id string) (VehicleInfo, Location, User, error)
}

type Presenter struct {
	controller *Controller
}

func NewPresenter(controller *Controller) *Presenter {
	return &Presenter{
		controller: controller,
	}
}

func (p *Presenter) GetVehicleByID(ctx *gin.Context) {

	vehicleID := ctx.Param("rentalID")

	if vehicleID == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("missing rentalID parametes"))
	}

	vehicleInfo, location, user, err := p.controller.GetVehicleByID(vehicleID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("cannot fetch a single rental by id: %w", err))
	}

	ctx.JSON(http.StatusOK, toGetRentalResponse(vehicleInfo, location, user))
}

func toGetRentalResponse(vehicleInfo VehicleInfo, location Location, user User) GetRentalResponse {
	return GetRentalResponse{
		vehicleInfo,
		location,
		user,
	}
}
