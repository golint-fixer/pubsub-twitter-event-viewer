package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/gorilla/websocket"
)

var subscription *pubsub.Subscription
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func configureSubscription() {
	client, err := pubsub.NewClient(appContext, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	subscription = client.Subscription(subscriptionName)
}

func subscribe(m chan []byte) {
	err := subscription.Receive(appContext, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		m <- msg.Data
	})
	if err != nil {
		log.Fatal(err)
	}
}

func wsWrite(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		log.Fatalf("Error while upgrading WS: %v", err)
	}

	for {
		select {
		case m := <-msgCh:
			if wErr := conn.WriteMessage(websocket.TextMessage, m); wErr != nil {
				log.Printf("Error on write message: %v", wErr)
			}
		}
	}

}
