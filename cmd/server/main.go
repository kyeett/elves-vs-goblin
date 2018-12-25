package main

import (
	"os"
	"os/signal"

	"github.com/kyeett/elves-vs-goblin/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	// world := world.NewWorld()
	quit := make(chan bool)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		logrus.Info("Received ctrl+C, shut down server")
		quit <- true
	}()

	s := server.NewDefault()
	s.Start(quit)

}
