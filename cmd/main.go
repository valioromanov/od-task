package main

import (
	"fmt"
	"net/http"
	"od-task/cmd/internal"
	"od-task/pkg/app"
	"od-task/pkg/repository/postgresql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleError(err error, m string) {
	if err != nil {
		app.Crash(fmt.Errorf("%s: %w", m, err))
	}
}

func main() {

	rentalRepo := postgresql.NewRentalRepository()
	presenter := internal.NewPresenter(internal.NewController(rentalRepo))
	handler := gin.New()

	handler.GET("/rental/:rentalID", presenter.GetVehicleByID)
	logrus.Info("starting http server...")
	httpServer := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			handleError(err, "server returned an error")
		}
	}()

	app.WaitExitSignal()
	logrus.Info("shutting down the application")
}
