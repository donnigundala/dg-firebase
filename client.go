package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"github.com/donnigundala/dg-firebase/fcm"
	"google.golang.org/api/option"
)

// Client wraps the Firebase App and provides access to services.
type Client struct {
	app *firebase.App
}

// NewClient creates a new Firebase client.
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	var opts []option.ClientOption
	if cfg.CredentialsJSON != "" {
		opts = append(opts, option.WithCredentialsJSON([]byte(cfg.CredentialsJSON)))
	} else if cfg.CredentialsFile != "" {
		opts = append(opts, option.WithCredentialsFile(cfg.CredentialsFile))
	}

	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	return &Client{app: app}, nil
}

// App returns the underlying Firebase App.
func (c *Client) App() *firebase.App {
	return c.app
}

// Auth returns a new Firebase Auth client.
func (c *Client) Auth(ctx context.Context) (*auth.Client, error) {
	return c.app.Auth(ctx)
}

// Firestore returns a new Firestore client.
func (c *Client) Firestore(ctx context.Context) (*firestore.Client, error) {
	return c.app.Firestore(ctx)
}

// Storage returns a new Firebase Storage client.
func (c *Client) Storage(ctx context.Context) (*storage.Client, error) {
	return c.app.Storage(ctx)
}

// FCM returns a new Firebase Cloud Messaging client.
func (c *Client) FCM(ctx context.Context) (*fcm.Client, error) {
	return fcm.NewClient(ctx, c.app)
}
