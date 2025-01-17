package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/anditakaesar/uwa-pubsub-init/internal/env"
	"github.com/anditakaesar/uwa-pubsub-init/xpubsub"
)

func main() {
	env.InitEnv()

	psClient := xpubsub.NewPubSubClient(&env.PubSubConfig{})

	topicSpec, err := xpubsub.LoadTopic(env.PubSubInitFile())
	if err != nil {
		slog.Error("failed to load topic", "error", err)
	}
	slog.Info(fmt.Sprintf("loading topics: %d", len(topicSpec)))

	err = psClient.CreateTopicAndSubscriptions(context.Background(), topicSpec)
	if err != nil {
		slog.Error("failed to create topic and subscriptions", "error", err)
	}
}
