package mercury

func (m *Mercury) Receiver(subject string, channel interface{}) error {
	sub, err := m.ec.BindRecvChan(subject, channel)
	if err != nil {
		return err
	}
	m.addSubscription(subject, sub)
	return nil
}

func (m *Mercury) Sender(subject string, channel interface{}) error {
	err := m.ec.BindSendChan(subject, channel)
	if err != nil {
		return err
	}
	return nil
}
