package internal_test

import (
	"fmt"
	"od-task/cmd/internal"
	"od-task/cmd/internal/mocks"
	"od-task/pkg/helper/mockutil"
	"od-task/pkg/repository/postgresql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	errMsg = "some-error"
)

var (
	rentalInfo = postgresql.FindResult{
		Rentals: postgresql.Rentals{
			ID:            5,
			UserForeignID: 1,
			Name:          "name",
			Type:          "type",
			Description:   "description",
			Sleeps:        5,
			PricePerDay:   16000,
			HomeCity:      "HomeCity",
			HomeZip:       "HomeZip",
			HomeState:     "HomeState",
		},
		Users: postgresql.Users{
			UserID:    1,
			FirstName: "FirstName",
			LastName:  "LastName",
		},
	}

	expectedResult = internal.GetRentalResponse{
		VehicleInfo: internal.VehicleInfo{
			VehicleID:   rentalInfo.ID,
			Name:        rentalInfo.Name,
			Description: rentalInfo.Description,
			Type:        rentalInfo.Type,
			Sleeps:      rentalInfo.Sleeps,
			Price: internal.Price{
				Day: rentalInfo.PricePerDay,
			},
		},
		Location: internal.Location{
			City:  rentalInfo.HomeCity,
			Zip:   rentalInfo.HomeZip,
			State: rentalInfo.HomeState,
		},
		User: internal.User{
			UserID:    rentalInfo.UserID,
			FirstName: rentalInfo.FirstName,
			LastName:  rentalInfo.LastName,
		},
	}
)

var _ = Describe("Controller", mockutil.Mockable(func(helper *mockutil.Helper) {
	var (
		mockRepo   *mocks.MockRentalRepository
		controller *internal.Controller
	)

	BeforeEach(func() {
		mockRepo = mocks.NewMockRentalRepository(helper.Controller())
		controller = internal.NewController(mockRepo)
	})

	Describe("GetRentalByID", func() {
		var id = "5"

		When("repository returns an error", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().FindById(id).Return(postgresql.FindResult{}, fmt.Errorf(errMsg))
			})

			It("should return an error", func() {
				_, err := controller.GetRentalByID(id)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(errMsg))
			})
		})

		When("repository returns a proper response", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().FindById(id).Return(rentalInfo, nil)
			})

			It("should succeed", func() {
				rental, err := controller.GetRentalByID(id)
				Expect(err).ToNot(HaveOccurred())
				Expect(rental).To(Equal(expectedResult))
			})
		})
	})

	Describe("GetFilteredRentals", func() {
		filters := map[string][]string{
			"limit": {"1"},
			"near":  {"11.54", "125.15"},
		}
		filteredDBResponse := []postgresql.FindResult{
			rentalInfo,
		}
		filteredResponse := []internal.GetRentalResponse{
			expectedResult,
		}
		When("repository returns an error", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().FindByFilters(filters).Return(nil, fmt.Errorf(errMsg))
			})

			It("should return an error", func() {
				_, err := controller.GetFilteredRentals(filters)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(errMsg))
			})
		})

		When("repository returns a proper response", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().FindByFilters(filters).Return(filteredDBResponse, nil)
			})

			It("should return an error", func() {
				filteredRental, err := controller.GetFilteredRentals(filters)
				Expect(err).ToNot(HaveOccurred())
				Expect(filteredRental).To(Equal(filteredResponse))
			})
		})
	})

}))
