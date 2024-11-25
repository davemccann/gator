package config

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

const (
	configFilename     = ".gatorconfig.json"
	defaultDatabaseURL = "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
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

func EnsureConfigExists() error {
	filepath, err := getConfigFilepath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}
		return err
	}

	defer file.Close()

	cfg := Config{
		DbURL:           defaultDatabaseURL,
		CurrentUserName: "",
	}

	bytes, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	if _, err := file.Write(bytes); err != nil {
		return err
	}

	return nil
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
