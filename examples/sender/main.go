package main

import (
	"fmt"
	"time"

	"github.com/nats-io/go-nats"

	"github.com/dashotv/mercury"
)

func main() {
	m, err := mercury.New("testing", nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("starting sender...")
	channel := make(chan *mercury.Message, 5)
	m.Sender("testing", channel)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("* sending message")
			channel <- &mercury.Message{Content: time.Now().String(), Sender: "sender"}
		}
	}()

	select {
	case <-time.After(30 * time.Second):
		fmt.Println("timeout")
	}
}
