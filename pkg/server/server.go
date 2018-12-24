package server

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/player"
	"github.com/kyeett/elves-vs-goblin/pkg/transport"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/elves-vs-goblin/pkg/world"
	"github.com/nats-io/nats"
)

type Server struct {
	world   *world.World
	conn    *nats.Conn
	encConn *nats.EncodedConn
}

var mutex = sync.RWMutex{}

var serverStartedTestHook = func() {}
var postActionTestHook = func() {}

type emptyCloser struct {
	close func() error
}

func (c emptyCloser) Close() error {
	c.close()
	return nil
}

func serverChannels(nc *nats.Conn) (chan *nats.Msg, chan *nats.Msg, io.Closer) {
	connectChan := make(chan *nats.Msg, 64)
	actionChan := make(chan *nats.Msg, 64)

	connectSub, err := nc.ChanSubscribe("connect", connectChan)
	if err != nil {
		log.Fatal(err)
	}

	actionSub, err := nc.ChanSubscribe("action", actionChan)
	if err != nil {
		log.Fatal(err)
	}

	f := func() error {
		connectSub.Unsubscribe()
		connectSub.Drain()

		actionSub.Unsubscribe()
		actionSub.Drain()
		return nil
	}
	closer := emptyCloser{
		close: f,
	}
	return connectChan, actionChan, closer
}

func NewDefaultServer() Server {
	nc, c, err := transport.ServerConnections()
	if err != nil {
		log.Fatal(err)
	}

	wrld := world.NewDefaultWorld()
	return Server{
		world:   &wrld,
		conn:    nc,
		encConn: c,
	}
}

func (s Server) Start(cancel <-chan bool) {
	log.Info("Starting server...")
	connectChan, actionChan, closer := serverChannels(s.conn)

	s.gameLoop(connectChan, actionChan, cancel)
	closer.Close()
	s.Shutdown()
}

func (s *Server) StartSendingState() {

	ticker := time.NewTicker(20 * time.Millisecond)
	for {
		<-ticker.C
		mutex.RLock()
		log.Info("Start sending state")
		s.encConn.Publish("state", s.world)
		mutex.RUnlock()
	}
}

func (s *Server) gameLoop(connectChan, actionChan chan *nats.Msg, cancel <-chan bool) {
	// Used in tests to wait for startup
	serverStartedTestHook()

	go s.StartSendingState()
	for {
		log.Info("Game Loop")
		select {
		case msg := <-connectChan:
			log.Info("Connect")
			s.handleConnect(msg)
		case msg := <-actionChan:
			log.Info("Action")
			s.handleAction(msg)
			postActionTestHook()
		case <-cancel:
			fmt.Println("Received an interrupt, cleanning up ...")
			return
		}
	}
}

func (s *Server) Shutdown() {
	s.conn.Close()
	s.encConn.Close()
}

func (s *Server) handleConnect(msg *nats.Msg) {
	p := player.NewDefaultPlayer()
	log.Infof("Player %s connected", p)
	s.world.AddPlayer(&p)
	log.Info(s.world)
	s.encConn.Publish(msg.Reply, &p)
}

func (s *Server) handleAction(msg *nats.Msg) {
	log.Info("Action performed")
}
