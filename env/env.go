package env

import (
	"os"
)

// Get environment variable
func Get(name string) string {
	return os.Getenv(name)
}

// GetWithDefault is to fetch environment variable, return defaultValue if env var is not exists
func GetWithDefault(name string, defaultValue string) string {
	e := Get(name)
	if e != "" {
		return e
	}

	return defaultValue
}
