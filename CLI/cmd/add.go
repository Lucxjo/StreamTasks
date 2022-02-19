package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Task string
)

var Add = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a task to the list",
	Long:  `Add a task to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("requires a name")
		} else {
			name := strings.Join(args, " ")

			_, err := http.PostForm("http://localhost:8080/add-task", url.Values{"task": {name}})
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Added: ", name)
			}
		};
	},
}