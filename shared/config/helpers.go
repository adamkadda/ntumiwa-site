package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getEnvInt(key string, defaultValue int) int {
	valueString := os.Getenv(key)
	if valueString == "" {
		fmt.Fprintf(os.Stdout, "[APIConfig] %s not set, using default: %d\n", key, defaultValue)
		return defaultValue
	}

	// Atoi(str) is equivalent to ParseInt(str, 10, 0)
	value, err := strconv.Atoi(valueString)
	if err != nil {
		fmt.Fprintf(os.Stdout, "[APIConfig] %s: %s invalid, using default: %v\n", key, valueString, defaultValue)
		return defaultValue
	}

	return value
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	valueString := os.Getenv(key)
	if valueString == "" {
		fmt.Fprintf(os.Stdout, "[APIConfig] %s not set, using default: %s\n", key, defaultValue)
		return defaultValue
	}

	value, err := time.ParseDuration(valueString)
	if err != nil {
		fmt.Fprintf(os.Stdout, "[APIConfig] %s: %s invalid, using default: %v\n", key, valueString, defaultValue)
		return defaultValue
	}

	return value
}
