package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/views"
	"github.com/kyeett/elves-vs-goblin/pkg/world"
)

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
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	world := world.NewWorld()
	view := views.NewView(&world)

	go world.Start()

	ticker := time.NewTicker(1000 * time.Millisecond)
	for range ticker.C {
		fmt.Println("Draw world!")
		fmt.Println(view)
		fmt.Println("🐳   🐙")
	}
}
