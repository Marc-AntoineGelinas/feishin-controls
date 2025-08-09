package cmd

import (
	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pauseCmd)
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause the current track",
	Long:  "Pause the currently selected track if it is currently playing",
	Run: func(cmd *cobra.Command, args []string) {
		controls.ClientSimpleEvent(controls.Pause)
	},
}
