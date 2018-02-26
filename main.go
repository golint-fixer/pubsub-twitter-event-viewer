package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"
)

var (
	appContext       context.Context
	projectID        string
	subscriptionName string
	serverPort       int
)

func main() {

	// config
	flag.IntVar(&serverPort, "port", 8080, "Server port")
	flag.StringVar(&projectID, "project", os.Getenv("GCLOUD_PROJECT"), "Google PubSub project ID")
	flag.StringVar(&subscriptionName, "subscription", "tweets-sub", "Google PubSub topic subscription [tweets-sub]")
	flag.Parse()

	if projectID == "" || subscriptionName == "" {
		log.Fatalf("Project and subscription arguments required: P:%s S:%s",
			projectID, subscriptionName)
	}

	// context
	ctx, cancel := context.WithCancel(context.Background())
	appContext = ctx

	go func() {
		// Wait for SIGINT and SIGTERM (HIT CTRL-C)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		log.Println(<-ch)
		cancel()
		os.Exit(0)
	}()

	// subscription
	configureSubscription()

	// init web server
	if err := startServer(serverPort); err != nil {
		log.Fatal(err)
		return
	}

}
