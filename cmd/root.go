package cmd

import (
	"fmt"

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
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Print("config.yml wasn't found in the running directory. Create and populate it.")
		} else {
			fmt.Print("config.yml was found, but something else went wrong.")
		}
	}
}
