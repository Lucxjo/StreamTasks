package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lucxjo/streamtasks/shared/models"
	"github.com/spf13/cobra"
)

var GetRoot = &cobra.Command{
	Use:   "get",
	Short: "Get tasks",
	Long:  `Get tasks`,
}

var GetAll = &cobra.Command{
	Use:   "all",
	Short: "Get all tasks",
	Long:  `Get all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		var decod []models.Task
		resp, _ := http.Get("http://localhost:8080/tasks")

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(body, &decod)
		
		if len(decod) != 0 {
				fmt.Println("ID", "    " + "Task")
				fmt.Println("--", "    " + "----")
			for _, task := range decod {
				fmt.Println(task.ID, "     " + task.Task)
			}
		} else {
			fmt.Println("No tasks found")
		}
	},
}

func RegisterGets() {
	GetRoot.AddCommand(GetAll)
}