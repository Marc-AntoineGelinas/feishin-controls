package cmd

import (
	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shuffleCmd)
}

var shuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Toggle shuffle",
	Long:  "Toggle feishin's shuffle between disabled and enabled",
	Run: func(cmd *cobra.Command, args []string) {
		controls.ClientSimpleEvent(controls.Shuffle)
	},
}
