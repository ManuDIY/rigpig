package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "rigpig",
}

func init() {
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(TestCmd)
	RootCmd.AddCommand(AgentCmd)
}
