package transport

import (
	"log"

	"github.com/nats-io/nats"
)

type Nats struct {
	conn *nats.EncodedConn
}

func DefaultNats() Nats {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	// Go type Publisher
	return Nats{
		conn: c,
	}
}

//Todo: how to handle closing of nc and c
func ServerConnections() (*nats.Conn, *nats.EncodedConn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, nil, err
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, nil, err
	}
	return nc, c, nil
}
