package cmd

import (
	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the current track",
	Long:  "Start playing the currently selected track if it is currently paused",
	Run: func(cmd *cobra.Command, args []string) {
		controls.ClientSimpleEvent(controls.Play)
	},
}
