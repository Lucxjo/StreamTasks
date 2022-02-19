package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "stasks",
	Short: "StreamTasks is a CLI for managing tasks",
	Long:  `StreamTasks is a CLI for managing tasks`,
}

func init() {
	rootCmd.AddCommand(Version)
	rootCmd.AddCommand(Add)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}