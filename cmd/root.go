package cmd

import (
	"fmt"
	"log"
	"os"
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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		configFilePath, err := config.GetConfigFilePath()
		if err != nil {
			log.Fatal("could not get config dir:", err)
		}

		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.AddConfigPath(filepath.Dir(configFilePath))

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				if cmd.Name() != "init" {
					fmt.Println("No config file found. Run the <init> command to generate it")
					os.Exit(1)
				}
			} else {
				log.Fatal("config.yml was found, but something else went wrong:\n", err)
			}
		}
		websocket.Authenticate()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
