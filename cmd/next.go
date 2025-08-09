package cmd

import (
	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nextCmd)
}

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Go to next track",
	Long:  "Skip to next track in the queue",
	Run: func(cmd *cobra.Command, args []string) {
		controls.ClientSimpleEvent(controls.Next)
	},
}
