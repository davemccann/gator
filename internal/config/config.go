package config

import (
	"encoding/json"
	"os"
	"path"
)

const (
	configFilename = ".gatorconfig.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return SaveConfig(cfg)
}

func getConfigFilepath() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filepath := path.Join(userHomeDir, configFilename)

	return filepath, nil
}

func ReadConfig() (Config, error) {

	filepath, err := getConfigFilepath()
	if err != nil {
		return Config{}, err
	}

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func SaveConfig(cfg *Config) error {
	filepath, err := getConfigFilepath()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath, bytes, 0644); err != nil {
		return err
	}

	return nil
}
