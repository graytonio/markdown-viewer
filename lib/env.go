package lib

import (
	"os"
)

type Config struct {
	Port   string
	MDRoot string
}

var config Config

func init() {
	config = Config{
		Port:   get_env_with_default("PORT", "9000"),
		MDRoot: get_env_with_default("MD_ROOT", "/markdown"),
	}
}

func GetConfig() Config {
	return config
}

func get_env_with_default(key string, def string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return def
	}
	return value
}
