package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "feishin-controls",
	Short: "feishin-controls is a cli tool to control feishin",
	Long:  "feishin-controls is a cli tool to control feishin via the websockets using the Remote Control",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
	initConfig()
	websocket.Authenticate()
}

func getConfigDir() (string, error) {
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

	return filepath.Join(configDir, "feishin-controls"), nil
}

func configFileExist(configFile string) bool {
	_, err := os.Stat(configFile)
	return err == nil
}

func createConfigFile(configPath string, configFile string) {
	err := os.MkdirAll(configPath, 0755)
	if err != nil {
		log.Fatal("could not create config file with permissions:", err)
	}

	viper.Set("URL", "localhost:4333")
	viper.Set("Username", "")
	viper.Set("Password", "")

	err = viper.SafeWriteConfigAs(configFile)
	if err != nil {
		log.Fatal("could not write config file:", err)
	}
}

func initConfig() {
	configPath, err := getConfigDir()
	if err != nil {
		log.Fatal("could not get config dir:", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(configPath)

	configFile := filepath.Join(configPath, "config.yml")

	if !configFileExist(configFile) {
		createConfigFile(configPath, configFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Print("config.yml wasn't found in the running directory. Create and populate it.")
		} else {
			fmt.Print("config.yml was found, but something else went wrong.")
		}
	}
}
