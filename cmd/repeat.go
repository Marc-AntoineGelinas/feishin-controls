package cmd

import (
	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(repeatCmd)
}

var repeatCmd = &cobra.Command{
	Use:   "repeat",
	Short: "Cycle repeat",
	Long:  "Cycle feishin's repeat between repeat none, all and one",
	Run: func(cmd *cobra.Command, args []string) {
		controls.ClientSimpleEvent(controls.Repeat)
	},
}
