package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config struct represent config file
type Config struct {
	Path string `json:"path"`
	File string `json:"file"`
}

// GetOrCreateConfig method get config or create default
func GetOrCreateConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := fmt.Sprintf("%s/.imgurLoader", homeDir)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, err := os.Create(configPath)
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf(`{"path": "%s/Pictures", "file": "\\.png"}`, homeDir))
		if err != nil {
			return nil, err
		}
	}

	fileData, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	var config Config

	errUm := json.Unmarshal(fileData, &config)

	if errUm != nil {
		println("Error while parsing config")
	}

	return &config, nil
}
