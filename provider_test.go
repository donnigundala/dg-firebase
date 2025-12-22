package dgfirebase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirebaseServiceProvider_Name(t *testing.T) {
	provider := &FirebaseServiceProvider{}
	assert.Equal(t, "firebase", provider.Name())
}

func TestFirebaseServiceProvider_Version(t *testing.T) {
	provider := &FirebaseServiceProvider{}
	assert.Equal(t, "1.1.0", provider.Version())
}

func TestFirebaseServiceProvider_Dependencies(t *testing.T) {
	provider := &FirebaseServiceProvider{}
	deps := provider.Dependencies()

	assert.NotNil(t, deps)
	assert.Empty(t, deps, "dg-firebase should have no dependencies")
}
