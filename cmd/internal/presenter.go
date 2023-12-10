package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen --source=presenter.go --destination mocks/presenter.go --package mocks
type RentalController interface {
	GetRentalByID(id string) (GetRentalResponse, error)
	GetFilteredRentals(filters map[string][]string) ([]GetRentalResponse, error)
}

type Presenter struct {
	Controller RentalController
}

func NewPresenter(controller RentalController) *Presenter {
	return &Presenter{
		Controller: controller,
	}
}

func (p *Presenter) GetSingleRentalByID(ctx *gin.Context) {
	vehicleID := ctx.Param("rentalID")
	if vehicleID == "" {
		ctx.JSON(http.StatusBadRequest, NewAPIError("missing rentalID parametes", http.StatusBadRequest))
		return
	}

	rental, err := p.Controller.GetRentalByID(vehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewAPIError(fmt.Errorf("cannot fetch a single rental by id: %w", err).Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, rental)
}

func (p *Presenter) GetRentalsByFilters(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	rentals, err := p.Controller.GetFilteredRentals(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewAPIError(fmt.Errorf("cannot fetch a filtered rentals: %s", err.Error()).Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, rentals)
}
