package transport

import (
	"log"

	"github.com/nats-io/nats"
)

type Messager interface {
	Send(v interface{})
}

type Nats struct {
	conn *nats.EncodedConn
}

func (n Nats) Send(v interface{}) {
	n.conn.Publish("player", &v)
}

func DefaultNats() Nats {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// // Simple Async Subscriber
	// nc.Subscribe("chat", func(m *nats.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// })

	// nc.Publish("chat", []byte("Hej"))

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	// defer c.Close()

	// Simple Publisher

	// Go type Subscriber
	// c.Subscribe("chat", func(p *person) {
	// 	fmt.Printf("Received a person: %+v\n", p)
	// })

	// Go type Publisher
	return Nats{
		conn: c,
	}
}

//Todo: how to handle closing of nc and c
func ServerConnections() (*nats.Conn, *nats.EncodedConn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, nil, err
	}
	return nc, c, nil
}
