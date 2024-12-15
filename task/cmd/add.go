package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/rajatnai/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		fmt.Printf("Task Created :%d\n", id)
		if err != nil {
			log.Fatal("Error in the opening file")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
