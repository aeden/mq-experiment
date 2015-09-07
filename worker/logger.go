// Log messages from nats
package main

import (
	"github.com/nats-io/nats"
	"log"
)

func runLogger() {
	log.Printf("Starting worker")

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

	subject := "cast"

	subscription, err := c.Subscribe(subject, func(s string) {
		log.Printf("Received a message: %v", s)
	})

	if err != nil {
		log.Printf("Failed to subscribe to subject %s: %v", subject, err)
		return
	}

	log.Printf("Listening for messages on %s", subscription.Subject)
}
