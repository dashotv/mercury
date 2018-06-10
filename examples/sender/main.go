package main

import (
	"fmt"
	"github.com/dashotv/mercury"
	"github.com/nats-io/go-nats"
	"time"
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
