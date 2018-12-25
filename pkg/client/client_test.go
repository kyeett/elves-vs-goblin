package client

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"

	"github.com/kyeett/elves-vs-goblin/pkg/server"
	"github.com/kyeett/elves-vs-goblin/pkg/world"
	"github.com/nats-io/nats"
	"github.com/pkg/errors"
)

func Test_move(t *testing.T) {
	s := server.NewDefaultServer()
	quit := make(chan bool)
	go s.Start(quit)

	// Wait for server to start up
	c := NewClient()

	retries := 3
	err := retryFunction(retries, 10*time.Millisecond, c.Connect)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "client failed to connect with %d attempts", retries))
	}

	// Receive updates
	stateChan := make(chan *nats.Msg, 64)
	_, _ = c.conn.ChanSubscribe("state", stateChan)

	timeout := 50 * time.Millisecond
	for i := 0; i < 2; i++ {
		c.Move(1, 2)
		select {
		case msg := <-stateChan:
			var wrld world.World
			err := json.Unmarshal(msg.Data, &wrld)
			if err != nil {
				t.Fatal(err)
			}
			c.world = &wrld
			// Todo: fix for multiplayer :-)
			c.Player.Coord = c.world.Players[0].Coord

		case <-time.After(timeout):
			t.Fatalf("\nDid not expected response within %s", timeout)
		}
	}

	expected := geom.Coord{X: 2, Y: 4}
	if c.Coord != expected {
		t.Fatalf("Expected %s, got %s", expected, c.Coord)
	}

	quit <- true
}

// func getClientPlayer(ID string, []player.Player) player.Player {
// 	for _, p := range players {

// 	}
// }

func retryFunction(retries int, delay time.Duration, f func() error) error {
	var err error
	for i := 0; i < retries; i++ {
		err = f()
		if err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return err
}
