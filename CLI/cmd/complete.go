package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var Complete = &cobra.Command{
	Use:   "complete [id]",
	Short: "Complete a task",
	Long:  `Complete a task`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("requires an id")
		} else {
			id := args[0]
			_, err := http.PostForm("http://localhost:8080/complete-task", url.Values{"id": {id}})
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Completed: ", id)
			}
		};
	},
}