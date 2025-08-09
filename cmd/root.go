package cmd

import (
	"log"
	"path/filepath"

	"github.com/marc-antoinegelinas/feishin-controls/internal/config"
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

func initConfig() {
	configFilePath, err := config.GetConfigFilePath()
	if err != nil {
		log.Fatal("could not get config dir:", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Dir(configFilePath))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config.yml wasn't found in the running directory. Create and populate it:\n", err)
		} else {
			log.Fatal("config.yml was found, but something else went wrong:\n", err)
		}
	}
}
