package dgfirebase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_CredentialsFile(t *testing.T) {
	cfg := Config{
		CredentialsFile: "/path/to/credentials.json",
	}

	assert.Equal(t, "/path/to/credentials.json", cfg.CredentialsFile)
	assert.Empty(t, cfg.CredentialsJSON)
}

func TestConfig_CredentialsJSON(t *testing.T) {
	jsonCreds := `{"type": "service_account", "project_id": "test"}`
	cfg := Config{
		CredentialsJSON: jsonCreds,
	}

	assert.Equal(t, jsonCreds, cfg.CredentialsJSON)
	assert.Empty(t, cfg.CredentialsFile)
}

func TestConfig_BothCredentials(t *testing.T) {
	// Test that both can be set (though only one should be used)
	cfg := Config{
		CredentialsFile: "/path/to/credentials.json",
		CredentialsJSON: `{"type": "service_account"}`,
	}

	assert.NotEmpty(t, cfg.CredentialsFile)
	assert.NotEmpty(t, cfg.CredentialsJSON)
}
