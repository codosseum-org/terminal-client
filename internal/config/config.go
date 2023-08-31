package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	modelconfig "github.com/codosseum-org/terminal-client/internal/model/config"
)

var (
	defaultConfig = modelconfig.Config {
		General: modelconfig.General{
			URL: "codosseum-tld.org",
			DefaultName: "username",
		},
	}
)

func GetConfig() (modelconfig.Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return modelconfig.Config{}, err
	}

	if !DoesConfigExist() {
		return modelconfig.Config{}, fmt.Errorf("configuration file does not exist")
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		return modelconfig.Config{}, err
	}

	var config modelconfig.Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		return modelconfig.Config{}, err
	}

	return config, nil
}

func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(configDir, "codosseum", "config.toml")

	return path, nil
}

func DoesConfigExist() bool {
	path, err := GetConfigPath()
	if err != nil {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func GenerateConfig() error {
	if DoesConfigExist() {
		return fmt.Errorf("configuration file already exists")
	}

	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(defaultConfig); err != nil {
		return err
	}

	return nil
}

func UpdateConfig(config modelconfig.Config) error {
	if !DoesConfigExist() {
		return fmt.Errorf("configuration file does not exist")
	}

	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(config); err != nil {
		return err
	}

	return nil
}
