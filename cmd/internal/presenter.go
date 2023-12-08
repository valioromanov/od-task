package internal

type RentalController interface {
	GetVehicleByID(id string) (VehicleInfo, Location, User)
}

type Presenter struct {
	controller Controller
}

func NewPresenter(controller Controller) *Presenter {
	return &Presenter{
		controller: controller,
	}
}

func (p *Presenter) GetVehicleByID(id string) GetRentalResponse {
	p.controller.GetVehicleByID(id)

	return GetRentalResponse{}
}
