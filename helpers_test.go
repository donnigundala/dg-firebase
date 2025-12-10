package firebase

import (
	"testing"

	"github.com/donnigundala/dg-core/foundation"
	"github.com/stretchr/testify/assert"
)

func TestResolve(t *testing.T) {
	app := foundation.New(".")
	/* cfg removed as unused/misconfigured for test */
	// We just need a dummy registration for testing resolution mechanics
	// In a real app, NewClient(ctx, cfg) would return a *Client.
	// Here we manually register a struct that matches *Client type (or just instantiate one if exported)

	// Create a dummy client to register. We don't need real connection
	client := &Client{}
	app.Instance("firebase", client)

	// Test Resolve
	c, err := Resolve(app)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestResolve_Error(t *testing.T) {
	app := foundation.New(".")

	// Test Resolve without registration
	c, err := Resolve(app)
	assert.Error(t, err)
	assert.Nil(t, c)
}

func TestMustResolve(t *testing.T) {
	app := foundation.New(".")
	app.Instance("firebase", &Client{})

	// Test MustResolve
	assert.NotPanics(t, func() {
		c := MustResolve(app)
		assert.NotNil(t, c)
	})
}

func TestMustResolve_Panic(t *testing.T) {
	app := foundation.New(".")

	// Test MustResolve panics without registration
	assert.Panics(t, func() {
		MustResolve(app)
	})
}

func TestInjectable(t *testing.T) {
	app := foundation.New(".")
	app.Instance("firebase", &Client{})

	inject := NewInjectable(app)

	assert.NotPanics(t, func() {
		c := inject.Firebase()
		assert.NotNil(t, c)
	})
}
