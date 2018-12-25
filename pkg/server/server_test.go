package server

import (
	"testing"

	"github.com/kyeett/elves-vs-goblin/pkg/client"
)

func Test_connect(t *testing.T) {
	s := NewDefaultServer()
	quit := make(chan bool)

	started := make(chan bool)
	serverStartedTestHook = func() {
		started <- true
	}
	go s.Start(quit)

	// Wait for server to start up
	<-started

	c := client.NewClient()
	err := c.Connect()
	if err != nil {
		t.Fatal(err)
	}
	// c.Move(1, 0)
	// c.Move(0, 1)

	// stateChan := c.StateChan()
	// go s.StartSendingState()

	// timeout := 100 * time.Millisecond
	// select {
	// case msg := <-stateChan:
	// 	var wrld world.World
	// 	err := json.Unmarshal(msg.Data, &wrld)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	log.Info("State updated", wrld)

	// case <-time.After(timeout):
	// 	t.Fatalf("Did not expected response within %s: %s", timeout, err)
	// }

	quit <- true
}
