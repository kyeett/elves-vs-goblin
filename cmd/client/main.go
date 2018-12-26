package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/client"
	"github.com/kyeett/elves-vs-goblin/pkg/input"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	c := client.New(os.Stdout)
	inputCh := make(chan input.Command)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			inputCh <- input.Command(rand.Intn(4))
		}
	}()

	go func() {
		<-signalCh
		logrus.Info("Received ctrl+C, shutting down client")
		cancel()
	}()

	c.Connect()
	if err := c.Run(inputCh, ctx); err != nil {
		log.Fatal(err)
	}
}
