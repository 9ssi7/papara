// Package papara provides functions for interacting with the Papara API
package papara

// Config represents the configuration for the Papara client.
type Config struct {
	// ApiKey is the Papara API key.
	ApiKey string `env:"PAPARA_API_KEY"`

	// ApiUrl is the base URL for the Papara API.
	ApiUrl string `env:"PAPARA_API_URL"`
}
