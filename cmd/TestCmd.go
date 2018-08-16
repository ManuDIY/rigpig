package cmd

import (
	"github.com/spf13/cobra"
	"rigpig/internal"
)

var TestCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		_ = internal.UpdateAlgos()
	},
}
