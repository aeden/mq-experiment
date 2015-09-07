// Respond to messages from nats

package main

import (
	"github.com/nats-io/nats"
	"log"
)

func runResponder() {
	log.Printf("Starting responder")

	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Printf("Failed to connect to message queue: %s", err)
		return
	}

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Printf("Failed to create encoded connection for message queue: %s", err)
		return
	}

	log.Printf("Connected")

	subject := "call"

	subscription, err := c.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf("Received a message: %v\n", msg)
		c.Publish(msg.Reply, "pong")
	})

	if err != nil {
		log.Printf("Failed to subscribe to subject %s: %v", subject, err)
		return
	}

	log.Printf("Responding to messages on %s", subscription.Subject)

}
