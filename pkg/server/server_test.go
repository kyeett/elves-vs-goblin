package server

import (
	"context"
	"os"
	"testing"

	"github.com/kyeett/elves-vs-goblin/pkg/client"
)

func Test_Connect(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := NewDefault()

	started := make(chan bool)
	serverStartedTestHook = func() {
		started <- true
	}
	go s.Run(ctx)

	// Wait for server to start up
	<-started

	c := client.New(os.Stdout)
	err := c.Connect()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ConnectMultiple(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := NewDefault()

	started := make(chan bool)
	serverStartedTestHook = func() {
		started <- true
	}
	go s.Run(ctx)

	// Wait for server to start up
	<-started

	wanted := 40
	for i := 0; i < wanted; i++ {
		c := client.New(os.Stdout)
		err := c.Connect()
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()
	}

	if len(s.world.Players) != wanted {
		t.Fatalf("Expected %d connected players, got %d", wanted, len(s.world.Players))
	}
}
