package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/kyeett/elves-vs-goblin/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		logrus.Info("Received ctrl+C, shut down server")
		cancel()
	}()

	s := server.NewDefault()
	s.Run(ctx)
}
