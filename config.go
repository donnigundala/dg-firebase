package firebase

// Config holds the configuration for the Firebase provider.
// Use either CredentialsFile OR CredentialsJSON, not both.
type Config struct {
	// CredentialsFile is the path to the service account JSON file.
	// Use this OR CredentialsJSON (not both).
	CredentialsFile string `mapstructure:"credentials_file"`

	// CredentialsJSON is the raw JSON content of the service account.
	// Useful for passing credentials via environment variables.
	// Use this OR CredentialsFile (not both).
	CredentialsJSON string `mapstructure:"credentials_json"`
}
