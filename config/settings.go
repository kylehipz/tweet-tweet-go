package config

import "os"

var DEV_SETTINGS = map[string]string {
	"port": os.Getenv("PORT"),
}

func GetSettings() map[string]string {
	// env := os.Getenv("ENV")

	return DEV_SETTINGS
}
