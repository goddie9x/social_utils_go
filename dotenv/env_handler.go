package dotenv

import "os"

func GetEnvOrDefaultValue(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
