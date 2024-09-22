package pubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func PubSubClient(projectId string) *pubsub.Client {
	ctx := context.Background()

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
