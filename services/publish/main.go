package main

import (
	"api/pkg/logging"
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"time"
)

const (
	ServiceName = "Scheduler"
	Version     = "0.0.1"
)

func main() {
	logging.New(ServiceName, Version)
	defer logging.Close()

	client, err := dapr.NewClient()
	if err != nil {
		logging.Fatal(err.Error())
	}

	defer client.Close()

	ctx := context.Background()
	i := 1
	for _ = range time.Tick(time.Second * 3) {
		fmt.Println(fmt.Sprintf("Run Number: %d", i))
		data := []byte(fmt.Sprintf("Hello %d", i))

		err = client.SaveState(ctx, "store", fmt.Sprintf("%d", i), data)
		if err != nil {
			logging.Error(err.Error())
		}

		err = client.PublishEvent(ctx, "pubsub", "pubsub", data)
		if err != nil {
			logging.Error(err.Error())
		}
		i++
	}
}
