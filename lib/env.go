package lib

import (
	"errors"
	"os"
)

type VaultType string

func (vt VaultType) IsValid() error {
	switch vt {
	case GitVault, LocalVault:
		return nil
	}
	return errors.New("invalid leave type using default")
}

const (
	GitVault   VaultType = "git"
	LocalVault VaultType = "local"
)

type Config struct {
	Port          string
	MDRoot        string
	Vault         VaultType
	UpdateRate    string
	GitURL        string
	GitUpdateRate string
}

var config Config

func init() {
	config = Config{
		Port:          get_env_with_default("PORT", "9000"),              // Default to port 9000
		MDRoot:        get_env_with_default("MD_ROOT", "/markdown"),      // Default to /markdown for where notes are stored
		Vault:         get_vault_type_with_default("VAULT", LocalVault),  // Default to local vault
		UpdateRate:    get_env_with_default("UPDATE_RATE", "30 * * * *"), // Default update every 30 mins
		GitURL:        os.Getenv("GIT_URL"),
		GitUpdateRate: get_env_with_default("GIT_UPDATE", "30 * * * *"),
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

func get_vault_type_with_default(key string, def VaultType) VaultType {
	value := os.Getenv(key)
	if len(value) == 0 || VaultType(value).IsValid() != nil {
		return def
	}

	return VaultType(value)
}
