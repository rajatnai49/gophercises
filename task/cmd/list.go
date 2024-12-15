package cmd

import (
	"fmt"

	"github.com/rajatnai/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all incomplete task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println("Error in the fetching tasks")
			return
		}
        fmt.Println("You have following tasks to do:")
		for i := 0; i < len(tasks); i++ {
            fmt.Printf("%d. %v\n", i+1, tasks[i].Value.Task)
		}
        fmt.Println()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
