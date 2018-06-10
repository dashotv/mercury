package mercury

import "github.com/nats-io/go-nats"

func (m *Mercury) addSubscription(channel string, subscription *nats.Subscription) {
	if m.subscriptions[channel] == nil {
		m.subscriptions[channel] = []*nats.Subscription{}
	}
	m.subscriptions[channel] = append(m.subscriptions[channel], subscription)
}
