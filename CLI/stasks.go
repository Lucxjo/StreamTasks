package main

import (
	"os"

	"github.com/lucxjo/streamtasks/cli/commands"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "add":
			commands.Add()
		case "list":
			commands.List()
		case "help":
			commands.Help()
		default:
			commands.Help()
		}
	} else {
		commands.Help()
	}
}
