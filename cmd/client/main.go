package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/client"
	"github.com/kyeett/elves-vs-goblin/pkg/input"
)

func main() {

	c := client.NewClient()

	inputCh := make(chan input.Command)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			inputCh <- input.Command(rand.Intn(4))
		}

	}()

	c.Connect()
	if err := c.Run(inputCh); err != nil {
		log.Fatal(err)
	}
}
