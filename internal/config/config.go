package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type ConfigFields struct {
	Url      string
	Username string
	Password string
}

func GetConfigFilePath() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		configDir = os.Getenv("APPDATA")
	case "darwin":
		configDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support")
	default:
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	if configDir == "" {
		return "", fmt.Errorf("could not determine config directory for os")
	}

	return filepath.Join(configDir, "feishin-controls", "config.yml"), nil
}

func CreateConfigFile(configs ConfigFields) {
	configFile, err := GetConfigFilePath()
	if err != nil {
		log.Fatal("could not get config file", err)
	}

	err = os.MkdirAll(filepath.Dir(configFile), 0755)
	if err != nil {
		log.Fatal("could not create config file with permissions:", err)
	}

	viper.Set("URL", configs.Url)
	viper.Set("Username", configs.Username)
	viper.Set("Password", configs.Password)

	err = viper.WriteConfigAs(configFile)
	if err != nil {
		log.Fatal("could not write config file:", err)
	}

	fmt.Printf("Config file created: %s\n", configFile)
}
