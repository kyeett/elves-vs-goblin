package main

import (
	"fmt"
	"log"

	"github.com/kyeett/elves-vs-goblin/pkg/views"
	"github.com/kyeett/elves-vs-goblin/pkg/world"

	"github.com/kyeett/elves-vs-goblin/pkg/player"

	"github.com/nats-io/nats"
)

func main() {
	world := world.NewWorld()
	view := views.NewView(&world)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// Simple Async Subscriber
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	connectedPlayers := make(map[string]bool)
	c.Subscribe("player", func(p *player.Player) {
		fmt.Printf("Received a message: %+v\n", p)

		if _, connected := connectedPlayers[p.ID]; !connected {
			fmt.Println("New player connected")
			connectedPlayers[p.ID] = true
			world.NewPlayer(p)
		}
	})

	quit := make(chan int)
	fmt.Println(view)
	<-quit

	// C

}
