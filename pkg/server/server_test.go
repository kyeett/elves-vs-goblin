package server

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/client"
	log "github.com/sirupsen/logrus"

	"github.com/kyeett/elves-vs-goblin/pkg/transport"

	"github.com/kyeett/elves-vs-goblin/pkg/world"
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

	var w world.World
	_, c, err := transport.ServerConnections()
	if err != nil {
		t.Fatal(err)
	}

	timeout := 10 * time.Millisecond
	err = c.Request("connect", "reqesterrerer", &w, timeout)
	if err != nil {
		t.Fatalf("Did not expected response within %s: %s", timeout, err)
	}

	quit <- true
}

func Test_move(t *testing.T) {
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
	c.Move(1, 0)
	c.Move(0, 1)

	stateChan := c.StateChan()
	go s.StartSendingState()

	timeout := 100 * time.Millisecond
	select {
	case msg := <-stateChan:
		var wrld world.World
		err := json.Unmarshal(msg.Data, &wrld)
		if err != nil {
			t.Fatal(err)
		}
		log.Info("State updated", wrld)

	case <-time.After(timeout):
		t.Fatalf("Did not expected response within %s: %s", timeout, err)
	}

	// Wait for server to receive
	// actionDone := make(chan bool)
	// postActionTestHook = func() {
	// 	actionDone <- true
	// }

	// <-actionDone
	// Terminate test
	quit <- true
}
