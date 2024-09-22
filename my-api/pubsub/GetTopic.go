package pubsub

import "cloud.google.com/go/pubsub"

func GetTopic(client pubsub.Client, id string) *pubsub.Topic {
	return client.Topic(id)
}
