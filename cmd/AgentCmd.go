package cmd

import "github.com/spf13/cobra"

var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Remote agent",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	AgentCmd.Flags().String("server", "127.0.0.1", "Address of RigPig server.")
}
