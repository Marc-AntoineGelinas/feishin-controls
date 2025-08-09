package cmd

import (
	"fmt"
	"log"

	"github.com/marc-antoinegelinas/feishin-controls/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [url] [username] [password]",
	Short: "Initialize feishin-controls",
	Long:  "Initializes feishin-controls config file",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.ConfigFields
		if len(args) == 3 {
			cfg.Url = args[0]
			cfg.Username = args[1]
			cfg.Password = args[2]
		} else if len(args) == 0 {
			fmt.Print("To initialize feishin-controls, enable Feishin's remote control server in Settings->General\n")
			fmt.Print("Enter the url. By default it should be localhost:4333, unless you're using a reverse proxy.\n")

			_, err := fmt.Scanln(&cfg.Url)
			if err != nil {
				log.Fatal("failed to scan value:", err)
			}

			fmt.Print("Enter the username.\n")
			_, err = fmt.Scanln(&cfg.Username)
			if err != nil {
				log.Fatal("failed to scan value:", err)
			}

			fmt.Print("Enter the password.\n")
			_, err = fmt.Scanln(&cfg.Password)
			if err != nil {
				log.Fatal("failed to scan value:", err)
			}
		} else {
			log.Fatal("init takes either 0 or 3 arguments ([url] [username] [password])")
		}

		config.CreateConfigFile(cfg)
	},
}
