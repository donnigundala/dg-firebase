package firebase

import (
	"fmt"

	"github.com/donnigundala/dg-core/contracts/foundation"
)

// Resolve resolves the Firebase client from the container.
func Resolve(app foundation.Application) (*Client, error) {
	instance, err := app.Make("firebase")
	if err != nil {
		return nil, err
	}
	return instance.(*Client), nil
}

// MustResolve resolves the Firebase client or panics.
func MustResolve(app foundation.Application) *Client {
	client, err := Resolve(app)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve firebase client: %v", err))
	}
	return client
}

// Injectable can be embedded in structs to provide easy access to Firebase.
type Injectable struct {
	app foundation.Application
}

// NewInjectable creates a new Injectable instance.
func NewInjectable(app foundation.Application) *Injectable {
	return &Injectable{
		app: app,
	}
}

// Firebase returns the resolved Firebase client.
// It panics if the client cannot be resolved.
func (i *Injectable) Firebase() *Client {
	return MustResolve(i.app)
}
