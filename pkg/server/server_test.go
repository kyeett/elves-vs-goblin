package server

import (
	"context"
	"os"
	"testing"

	"github.com/kyeett/elves-vs-goblin/pkg/client"
)

func Test_connect(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := NewDefault()

	started := make(chan bool)
	serverStartedTestHook = func() {
		started <- true
	}
	go s.Start(ctx)

	// Wait for server to start up
	<-started

	c := client.New(os.Stdout)
	err := c.Connect()
	if err != nil {
		t.Fatal(err)
	}
	c.Run()

}
