package mercury

import (
	"github.com/nats-io/nats.go"
)

type Mercury struct {
	Name string
	URL  string
	nc   *nats.Conn
	ec   *nats.EncodedConn

	subscriptions map[string][]*nats.Subscription
}

type Handler func(message interface{})

type Message struct {
	Content string
	Sender  string
}

func New(name string, URL string) (*Mercury, error) {
	m := &Mercury{Name: name, URL: URL}

	nc, err := nats.Connect(URL)
	if err != nil {
		return nil, err
	}

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}

	m.subscriptions = map[string][]*nats.Subscription{}
	m.nc = nc
	m.ec = c
	return m, nil
}

func (m *Mercury) Status() nats.Status {
	return m.nc.Status()
}
