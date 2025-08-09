package cmd

import (
	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(previousCmd)
}

var previousCmd = &cobra.Command{
	Use:   "previous",
	Short: "Go to previous track or restart track",
	Long:  "Return to previous track in the queue or rewind track to start",
	Run: func(cmd *cobra.Command, args []string) {
		controls.ClientSimpleEvent(controls.Previous)
	},
}
