package cmd

import (
	"github.com/spf13/cobra"
	"rigpig/server"
)

var ServerCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		var server = server.NewServer()
		server.Start()
	},
}

func init() {

}
