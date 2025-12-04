package firebase

import (
	"context"

	"github.com/donnigundala/dg-core/contracts/foundation"
)

// FirebaseServiceProvider implements the PluginProvider interface.
type FirebaseServiceProvider struct {
	Config Config `config:"firebase"`
}

// Name returns the name of the plugin.
func (p *FirebaseServiceProvider) Name() string {
	return "dg-firebase"
}

// Version returns the version of the plugin.
func (p *FirebaseServiceProvider) Version() string {
	return "1.0.0"
}

// Dependencies returns the list of dependencies.
func (p *FirebaseServiceProvider) Dependencies() []string {
	return []string{}
}

// Register registers the Firebase service provider.
func (p *FirebaseServiceProvider) Register(app foundation.Application) error {
	// Config is auto-injected by the application if the struct tag `config` is present.
	// or we can rely on manual injection if auto-injection isn't fully set up for this pattern yet.
	// But based on the plan, auto-injection is supported.
	// Let's assume auto-injection works for now, or use container resolution if needed.
	// However, to be safe and explicit (and since I can't see InjectProviderConfig implementation yet),
	// I will stick to manual injection if I can resolve config, OR trust the auto-injection.

	// Actually, looking at the previous file view of application.go:
	// func (app *Application) Register(provider foundation.ServiceProvider) error {
	//     if err := InjectProviderConfig(provider); err != nil { ... }
	// }
	// So auto-injection IS called.

	// Register the client as a singleton
	app.Singleton("firebase", func() (interface{}, error) {
		return NewClient(context.Background(), p.Config)
	})

	return nil
}

// Boot boots the Firebase service provider.
func (p *FirebaseServiceProvider) Boot(app foundation.Application) error {
	// Firebase will be resolved when needed
	// No need to verify resolution here to avoid deadlock
	return nil
}
