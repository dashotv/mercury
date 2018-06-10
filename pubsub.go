package mercury

import (
	"fmt"
)

func (m *Mercury) Publish(channel string, message interface{}) error {
	err := m.ec.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mercury) Subscribe(channel string, f Handler) error {
	sub, err := m.ec.Subscribe(channel, f)
	if err != nil {
		return err
	}
	m.addSubscription(channel, sub)
	return nil
}

func (m *Mercury) Unsubscribe(channel string) error {
	var err error
	count := 0
	if len(m.subscriptions[channel]) == 0 {
		return nil
	}
	for _, sub := range m.subscriptions[channel] {
		err = sub.Unsubscribe()
		if err != nil {
			count++
		}
	}
	if count > 0 {
		return fmt.Errorf("error while unsubscribing, most recent: %s", err)
	}
	return nil
}
