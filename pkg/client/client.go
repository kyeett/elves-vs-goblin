package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/input"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/elves-vs-goblin/pkg/views"
	"github.com/kyeett/elves-vs-goblin/pkg/world"

	"github.com/kyeett/elves-vs-goblin/pkg/actions"
	"github.com/kyeett/elves-vs-goblin/pkg/transport"
	"github.com/nats-io/nats"

	"github.com/kyeett/elves-vs-goblin/pkg/player"
	"github.com/pkg/errors"
)

type Client struct {
	*player.Player
	world   *world.World
	conn    *nats.Conn
	encConn *nats.EncodedConn
	view    views.View
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
		view:    views.NewView(),
	}
}

func (c *Client) Connect() error {
	var p player.Player
	err := c.encConn.Request("connect", "I want to connect to the server", &c.Player, connectionTimeout)
	if err != nil {
		return errors.Wrap(err, "client failed to connect")
	}

	if p.ID == "" {
		return errors.Wrap(err, "client received unexpected response. No ID")
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

func (c *Client) Run(inputCh <-chan input.Command) error {

	stateChan := make(chan *nats.Msg, 64)
	sub, err := c.conn.ChanSubscribe("state", stateChan)
	if err != nil {
		return errors.Wrap(err, "client")
	}

	for {
		select {
		case cmd := <-inputCh:
			log.Infof("%s", input.Command(cmd))
			c.handleInput(cmd)
		case msg := <-stateChan:
			var wrld world.World
			err := json.Unmarshal(msg.Data, &wrld)
			if err != nil {
				return err
			}
			c.world = &wrld
			// Todo: fix for multiplayer :-)
			c.Player.Coord = c.world.Players[0].Coord
			fmt.Println(c.view.Draw(&wrld))

			// case <-cancel:
			// 	break
		}
	}

	sub.Unsubscribe()
	sub.Drain()
	return nil
}

func (c *Client) handleInput(cmd input.Command) {
	switch cmd {
	case input.MoveUp:
		c.Move(0, -1)
	case input.MoveDown:
		c.Move(0, 1)
	case input.MoveRight:
		c.Move(1, 0)
	case input.MoveLeft:
		c.Move(-1, 0)
	default:
		log.Errorf("unknown input %s. ignorning.", cmd)
	}
}
