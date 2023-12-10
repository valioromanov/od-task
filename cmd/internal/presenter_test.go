package internal_test

import (
	"od-task/cmd/internal"
	"od-task/pkg/helper/mockutil"

	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Presenter", mockutil.Mockable(func(helper *mockutil.Helper) {
	var (
		// recorder   *httptest.ResponseRecorder
		controller *internal.Controller
		// mockContext *gin.Context
		presenter *internal.Presenter
	)

	BeforeEach(func() {
		// recorder = httptest.NewRecorder()
		// mockContext, _ = gin.CreateTestContext(recorder)
		controller = internal.NewController(helper.Controller())
		presenter = internal.NewPresenter(controller)
	})
}))
