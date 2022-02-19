package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of StreamTasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("StreamTasks v0.0.0")
	},
}