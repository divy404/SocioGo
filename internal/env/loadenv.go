package env

import (
	"log"
	"os"
	"strconv"
)

func GetStringEnv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0] // ✅ use fallback
		}
		log.Fatalf("Env variable %s is required but not set", key)
	}
	return value
}

func GetIntEnv(key string, defaultValue ...int) int {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0] // ✅ use fallback
		}
		log.Fatalf("Env variable %s is required but not set", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0] // ✅ fallback if parsing fails
		}
		log.Fatalf("Unable to parse %s to int for %s", value, key)
	}
	return intValue
}
