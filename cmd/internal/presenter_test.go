package internal_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"od-task/cmd/internal"
	"od-task/cmd/internal/mocks"
	"od-task/pkg/helper/mockutil"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const errorMessage = "some-error"

var (
	expectedControllerResult = internal.GetRentalResponse{
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

var _ = Describe("Presenter", mockutil.Mockable(func(helper *mockutil.Helper) {
	var (
		recorder    *httptest.ResponseRecorder
		controller  *mocks.MockRentalController
		mockContext *gin.Context
		presenter   *internal.Presenter
	)

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		mockContext, _ = gin.CreateTestContext(recorder)
		controller = mocks.NewMockRentalController(helper.Controller())
		presenter = internal.NewPresenter(controller)
	})

	Describe("GetSingleRentalByID", func() {
		id := "1"
		var url = "http://abc.com/rental/"
		When("path param rentalID is missing", func() {
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("GET", url, nil)
			})

			It("should return StatusBadRequest with error body", func() {
				var errResp internal.APIError
				presenter.GetSingleRentalByID(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &errResp)).To(Succeed())
				Expect(errResp.Code).To(Equal(http.StatusBadRequest))
				Expect(errResp.Messgage).To(ContainSubstring("missing rentalID parametes"))
			})
		})

		When("Controller returns an error", func() {
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("GET", url, nil)
				controller.EXPECT().GetRentalByID(id).Return(internal.GetRentalResponse{}, fmt.Errorf(errorMessage))
			})

			It("should return StatusInternalServerError with error body", func() {
				mockContext.Params = []gin.Param{
					{Key: "rentalID", Value: id},
				}
				var errResp internal.APIError
				presenter.GetSingleRentalByID(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusInternalServerError))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &errResp)).To(Succeed())
				Expect(errResp.Code).To(Equal(http.StatusInternalServerError))
				Expect(errResp.Messgage).To(ContainSubstring("cannot fetch a single rental by id"))
			})
		})

		When("Controller returns a proper response", func() {
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("GET", url, nil)
				controller.EXPECT().GetRentalByID(id).Return(expectedControllerResult, nil)
			})

			It("should return StatusOK with rental body", func() {
				mockContext.Params = []gin.Param{
					{Key: "rentalID", Value: id},
				}
				var rentalResp internal.GetRentalResponse
				presenter.GetSingleRentalByID(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &rentalResp)).To(Succeed())
				Expect(rentalResp).To(Equal(expectedControllerResult))
			})
		})
	})

	FDescribe("GetRentalsByFilters", func() {
		var urlString = "http://abc.com/rental/"
		When("Controller returns an error", func() {
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("GET", urlString, nil)
				controller.EXPECT().GetFilteredRentals(gomock.Any()).Return(nil, fmt.Errorf(errorMessage))
			})

			It("should return StatusInternalServerError", func() {
				mockContext.Request.URL.RawQuery = "limit=1&near=\"11.54,125.15\""
				presenter.GetRentalsByFilters(mockContext)
				var errResp internal.APIError
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusInternalServerError))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &errResp)).To(Succeed())
				Expect(errResp.Code).To(Equal(http.StatusInternalServerError))
				Expect(errResp.Messgage).To(ContainSubstring("cannot fetch a filtered rentals"))
			})
		})

		When("Controller returns proper response", func() {
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("GET", urlString, nil)
				controller.EXPECT().GetFilteredRentals(gomock.Any()).Return([]internal.GetRentalResponse{expectedControllerResult}, nil)
			})

			It("should retrun a StatusOK with the filtered rental", func() {
				mockContext.Request.URL.RawQuery = "limit=1&near=\"11.54,125.15\""
				presenter.GetRentalsByFilters(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
				var resp []internal.GetRentalResponse
				Expect(json.Unmarshal(recorder.Body.Bytes(), &resp)).To(Succeed())
				Expect(len(resp)).To(Equal(1))
				Expect(resp[0]).To(Equal(expectedControllerResult))
			})
		})
	})
}))
