package main

import (
	"log"
	"os"
	"rigpig/cmd"
	"rigpig/internal"
)

func main() {
	f, _ := os.OpenFile(internal.Logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	log.SetOutput(f)

	cmd.RootCmd.Execute()
}
