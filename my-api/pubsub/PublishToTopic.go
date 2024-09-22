package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func PublishToTopic(topic *pubsub.Topic, message Message) (*pubsub.PublishResult, error) {
	ctx := context.Background()

	bytes, err := json.Marshal(message)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %v", err)
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(bytes),
	})

	return res, nil
}
