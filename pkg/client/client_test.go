package client

import (
	"context"
	"os"
	"testing"
	"testing/iotest"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/input"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"

	"github.com/kyeett/elves-vs-goblin/pkg/server"
	"github.com/pkg/errors"
)

func Test_move(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := server.NewDefault()
	go s.Run(ctx)

	c := New(iotest.TruncateWriter(os.Stdout, 0))
	retries := 3
	err := retryFunction(retries, 10*time.Millisecond, c.Connect)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "client failed to connect with %d attempts", retries))
	}

	updated := make(chan bool)
	postStateChangedHook = func() {
		updated <- true
	}

	inputCh := make(chan input.Command)
	go c.Run(ctx, inputCh)

	moves := []input.Command{
		input.MoveDown,
		input.MoveDown,
		input.MoveDown,
		input.MoveDown,
		input.MoveRight,
		input.MoveRight,
	}

	timeout := 50 * time.Millisecond
	for _, move := range moves {
		inputCh <- move
		select {
		case <-updated:
			// Do nothing
		case <-time.After(timeout):
			t.Fatalf("Sending inputs timed out after %s", timeout)
		}
	}

	expected := geom.Coord{X: 2, Y: 4}
	if c.Coord != expected {
		t.Fatalf("Expected %s, got %s", expected, c.Coord)
	}
}

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
