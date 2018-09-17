package gosample

import (
	"os"
)

// GetEnv environment variable
func GetEnv(name string) string {
	return os.Getenv(name)
}

// GetEnvWithDefault is to fetch environment variable, return defaultValue if env var is not exists
func GetEnvWithDefault(name string, defaultValue string) string {
	e := GetEnv(name)
	if e != "" {
		return e
	}

	return defaultValue
}
