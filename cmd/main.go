package main

import (
	"fmt"
	"net/http"
	"od-task/cmd/env"
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

	config, err := env.LoadAppConfig()

	handleError(err, "dailed to load app config")

	fmt.Println("config: ", config)

	rentalRepo := postgresql.NewRentalRepository(config)
	presenter := internal.NewPresenter(internal.NewController(rentalRepo))
	handler := gin.New()

	handler.GET("/rental/:rentalID", presenter.GetVehicleByID)
	handler.GET("/rental", presenter.GetFilteredVehicles)
	logrus.Info("starting http server...")
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
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
