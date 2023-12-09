package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func WaitExitSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sigChan
	signal.Stop(sigChan)
}

func Crash(err error) {
	logrus.WithError(err).Fatal("application has crashed")
}
