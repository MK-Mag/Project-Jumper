package Config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	IDE string `Json:"ide"`
}

func configpath() (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")

	if localAppData == "" {
		return "", fmt.Errorf("LOCALAPPDATA environment variable is not set")
	}
	return filepath.Join(localAppData, "PJ", "config.json"), nil
}

func ConfigLoad() (Config, error) {

	path, err := configpath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{IDE: "code"}, nil
		}
		return Config{}, fmt.Errorf("Read config file : %w ", err)
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("decode config file: %w", err)
	}

	return config, nil

}

func SaveConfig(config Config) error {

	path, err := configpath()
	if err != nil {
		return fmt.Errorf("get config path: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("write config file: %w", err)
	}

	return nil
}
