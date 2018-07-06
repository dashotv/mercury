package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/go-nats"

	"github.com/dashotv/mercury"
)

func main() {
	m, err := mercury.New("testing", nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("starting receiver...")
	channel := make(chan *mercury.Message, 5)
	m.Receiver("testing", channel)

	for {
		select {
		case r := <-channel:
			fmt.Printf("received: %#v\n", r)
		case <-time.After(30 * time.Second):
			fmt.Println("timeout")
			os.Exit(0)
		}
	}
}
