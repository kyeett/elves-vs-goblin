package client

import (
	"context"
	"encoding/json"
	"io"
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

var postStateChangedHook = func() {}

const connectionTimeout = 10 * time.Millisecond

func New(w io.Writer) Client {
	return Client{
		view: views.New(w),
	}
}

func (c *Client) Connect() error {
	if c.conn != nil || c.encConn != nil {
		return errors.New("client already connected")
	}

	conn, encConn, err := transport.ServerConnections()
	if err != nil {
		return errors.Wrap(err, "client")
	}
	c.conn, c.encConn = conn, encConn

	var p player.Player
	err = c.encConn.Request("connect", "I want to connect to the server", &c.Player, connectionTimeout)
	if err != nil {
		return errors.Wrap(err, "client failed to connect")
	}

	// if p.ID == "" {
	// 	return errors.Wrap(err, "client received unexpected response. No ID")
	// }
	c.Player = &p
	return nil
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}

	if c.encConn != nil {
		c.encConn.Close()
	}
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

func (c *Client) Run(ctx context.Context, inputCh <-chan input.Command) error {

	stateChan := make(chan *nats.Msg, 64)
	sub, err := c.conn.ChanSubscribe("state", stateChan)
	if err != nil {
		return errors.Wrap(err, "client")
	}
	defer sub.Unsubscribe()
	defer sub.Drain()

	for {
		select {
		case cmd := <-inputCh:
			c.handleInput(cmd)
		case msg := <-stateChan:

			// Emtpy queue
			for len(stateChan) > 0 {
				msg = <-stateChan
			}

			// TOdo: use function
			// c.handleStateChange(msg.Data)

			var wrld world.World
			err := json.Unmarshal(msg.Data, &wrld)
			if err != nil {
				return err
			}
			c.world = &wrld

			p, err := c.getPlayer(c.ID)
			if err != nil {
				return err
			}
			c.Player.Coord = p.Coord

			c.view.Draw(&wrld)

			postStateChangedHook()
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *Client) handleStateChange(msg *nats.Msg) error {
	var wrld world.World
	err := json.Unmarshal(msg.Data, &wrld)
	if err != nil {
		return err
	}
	c.world = &wrld

	p, err := c.getPlayer(c.ID)
	if err != nil {
		return err
	}
	c.Player.Coord = p.Coord
	return nil
}

func (c *Client) handleInput(cmd input.Command) {
	log.Debugf("client: handleInput: %s", cmd)
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

func (c *Client) getPlayer(ID string) (*player.Player, error) {
	for _, p := range c.world.Players {
		return p, nil
	}

	return nil, errors.New("invalid ID")
}

func (c *Client) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func (c *Client) SetLevel(level log.Level) {
	log.SetLevel(level)
}
