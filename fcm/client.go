package fcm

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

// Client wraps the Firebase Cloud Messaging client.
type Client struct {
	messaging *messaging.Client
	obs       *observability
}

// NewClient creates a new FCM client from a Firebase app.
func NewClient(ctx context.Context, app *firebase.App) (*Client, error) {
	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create messaging client: %w", err)
	}

	return &Client{
		messaging: messagingClient,
		obs:       newObservability(),
	}, nil
}

// Send sends a message to a single device.
func (c *Client) Send(ctx context.Context, message *messaging.Message) (string, error) {
	start := time.Now()
	res, err := c.messaging.Send(ctx, message)
	c.obs.record(ctx, "send", "single", start, err, 1)
	return res, err
}

// SendMulticast sends a message to multiple devices.
func (c *Client) SendMulticast(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error) {
	start := time.Now()
	res, err := c.messaging.SendMulticast(ctx, message)
	count := int64(0)
	if message != nil {
		count = int64(len(message.Tokens))
	}
	c.obs.record(ctx, "send_multicast", "multicast", start, err, count)
	return res, err
}

// SendAll sends multiple messages in a batch.
func (c *Client) SendAll(ctx context.Context, messages []*messaging.Message) (*messaging.BatchResponse, error) {
	start := time.Now()
	res, err := c.messaging.SendAll(ctx, messages)
	c.obs.record(ctx, "send_all", "batch", start, err, int64(len(messages)))
	return res, err
}

// SubscribeToTopic subscribes tokens to a topic.
func (c *Client) SubscribeToTopic(ctx context.Context, tokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return c.messaging.SubscribeToTopic(ctx, tokens, topic)
}

// UnsubscribeFromTopic unsubscribes tokens from a topic.
func (c *Client) UnsubscribeFromTopic(ctx context.Context, tokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return c.messaging.UnsubscribeFromTopic(ctx, tokens, topic)
}

// Client returns the underlying Firebase Messaging client for advanced usage.
func (c *Client) Client() *messaging.Client {
	return c.messaging
}
