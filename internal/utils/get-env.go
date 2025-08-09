package utils

import "os"

// GetEnv is used to getting Environments setted to microservice
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
