package main

import (
	"fmt"
	"log"

	"github.com/kyeett/elves-vs-goblin/pkg/views"
	"github.com/kyeett/elves-vs-goblin/pkg/world"

	"gonum.org/v1/gonum/mat"

	"github.com/nats-io/nats"
)

type World struct {
	*mat.Dense
}

const (
	Whale float64 = 1

	Tree        float64 = 20
	RoundedTree float64 = 21
)

const (
	width  = 8
	height = 8
)

func main() {

	world := world.NewDefaultWorld()
	view := views.NewView(&world)
	fmt.Println(view)

	fmt.Println("🐳   🐙")
}

func natsStuff() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// Simple Async Subscriber
	nc.Subscribe("chat", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	nc.Publish("chat", []byte("Hej"))

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// Simple Publisher
	c.Publish("chat", "Hello World")

	// EncodedConn can Publish any raw Go type using the registered Encoder
	type person struct {
		Name    string
		Address string
		Age     int
	}

	// Go type Subscriber
	c.Subscribe("chat", func(p *person) {
		fmt.Printf("Received a person: %+v\n", p)
	})

	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}

	// Go type Publisher
	c.Publish("chat", me)
}