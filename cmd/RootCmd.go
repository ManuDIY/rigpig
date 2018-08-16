package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "lucidminer",
}

func init() {
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(TestCmd)
}
