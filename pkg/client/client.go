package client

import (
	"log"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/actions"
	"github.com/kyeett/elves-vs-goblin/pkg/transport"
	"github.com/nats-io/nats"

	"github.com/kyeett/elves-vs-goblin/pkg/player"
	"github.com/pkg/errors"
)

type Client struct {
	*player.Player
	conn    *nats.Conn
	encConn *nats.EncodedConn
}

const connectionTimeout = 10 * time.Millisecond

func NewClient() Client {
	conn, encConn, err := transport.ServerConnections()
	if err != nil {
		log.Fatal(err)
	}

	return Client{
		conn:    conn,
		encConn: encConn,
	}
}

func (c *Client) Connect() error {
	var p player.Player
	err := c.encConn.Request("connect", "I want to connect to the server", &c.Player, connectionTimeout)
	if err != nil {
		return errors.Wrap(err, "client failed to connect")
	}

	c.Player = &p
	return nil
}

func (c *Client) Move(x, y int) error {
	if c.Player == nil {
		return errors.New("client has not connected to server. Use Connect first")
	}

	a := actions.Signal{
		ID:     c.ID,
		Coord:  c.Coord.Add(x, y),
		Action: actions.Move,
	}
	c.encConn.Publish("action", &a)
	return nil
}

func (c *Client) Run() error {
	c.Connect()

	stateChan := make(chan *nats.Msg, 64)
	sub, err := c.conn.ChanSubscribe("state", stateChan)
	if err != nil {
		return errors.Wrap(err, "client")
	}

	// select {
	// case msg := <-stateChan:
	// 	log.Info("State updated")
	// case  <- time.After(100*time.Millisecond)

	sub.Unsubscribe()
	sub.Drain()
	return nil
}

//Todo: fix closing of sub
func (c *Client) StateChan() chan *nats.Msg {
	stateChan := make(chan *nats.Msg, 64)
	_, _ = c.conn.ChanSubscribe("state", stateChan)
	return stateChan
}
