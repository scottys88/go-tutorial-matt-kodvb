package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func CreateSubscription(projectID, subID string, topic *pubsub.Topic) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	// topic of type https://godoc.org/cloud.google.com/go/pubsub#Topic
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	subscription := client.Subscription(subID)
	exists, subscriptionExistsError := subscription.Exists(ctx)
	if subscriptionExistsError != nil {
		return fmt.Errorf("pubsub.NewClient: %w", subscriptionExistsError)
	}

	if exists {
		fmt.Printf("Subscription already exists: %v", subID)
		return nil
	}

	sub, err := client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("CreateSubscription: %w", err)
	}
	fmt.Printf("Created subscription: %v\n", sub)
	return nil
}
