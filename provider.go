package dgfirebase

import (
	"context"

	"github.com/donnigundala/dg-core/contracts/foundation"
)

// FirebaseServiceProvider implements the PluginProvider interface.
type FirebaseServiceProvider struct {
	Config Config `config:"firebase"`
}

// NewFirebaseServiceProvider creates a new Firebase service provider.
func NewFirebaseServiceProvider() *FirebaseServiceProvider {
	return &FirebaseServiceProvider{}
}

// Name returns the name of the plugin.
func (p *FirebaseServiceProvider) Name() string {
	return Binding
}

// Version returns the version of the plugin.
func (p *FirebaseServiceProvider) Version() string {
	return Version
}

// Dependencies returns the list of dependencies.
func (p *FirebaseServiceProvider) Dependencies() []string {
	return []string{}
}

// Register registers the Firebase service provider.
func (p *FirebaseServiceProvider) Register(app foundation.Application) error {
	app.Singleton(Binding, func() (interface{}, error) {
		// Create client
		// We use context.Background() as this is a long-lived service
		client, err := NewClient(context.Background(), p.Config)
		if err != nil {
			return nil, err
		}
		return client, nil
	})

	return nil
}

// Boot boots the Firebase service provider.
func (p *FirebaseServiceProvider) Boot(app foundation.Application) error {
	// Firebase will be resolved when needed
	// No need to verify resolution here to avoid deadlock
	return nil
}

// Shutdown gracefully stops the Firebase service.
func (p *FirebaseServiceProvider) Shutdown(app foundation.Application) error {
	// Firebase App doesn't require explicit shutdown generally,
	// but we implement this to satisfy the ShutdownProvider interface
	// and remain consistent with other plugins.
	return nil
}
