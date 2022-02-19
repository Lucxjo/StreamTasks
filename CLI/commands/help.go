package commands

import "fmt"

func Help() {
	Header()

	fmt.Println("StreamTasks is a tool for managing tasks in a stream.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("  stasks [command] [options]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("  add	 Add a task to the widget")
}