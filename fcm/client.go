package fcm

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

// Client wraps the Firebase Cloud Messaging client.
type Client struct {
	messaging *messaging.Client
}

// NewClient creates a new FCM client from a Firebase app.
func NewClient(ctx context.Context, app *firebase.App) (*Client, error) {
	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create messaging client: %w", err)
	}

	return &Client{
		messaging: messagingClient,
	}, nil
}

// Send sends a message to a single device.
func (c *Client) Send(ctx context.Context, message *messaging.Message) (string, error) {
	return c.messaging.Send(ctx, message)
}

// SendMulticast sends a message to multiple devices.
func (c *Client) SendMulticast(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error) {
	return c.messaging.SendMulticast(ctx, message)
}

// SendAll sends multiple messages in a batch.
func (c *Client) SendAll(ctx context.Context, messages []*messaging.Message) (*messaging.BatchResponse, error) {
	return c.messaging.SendAll(ctx, messages)
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
