package cmd

import (
	"strconv"
	"github.com/rajatnai/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark the task on your TODO list completed",
	Run: func(cmd *cobra.Command, args []string) {
		var input []int
		for _, v := range args {
			i, err := strconv.Atoi(v)
			if err == nil {
				input = append(input, i)
			}
		}
		db.CompleteTask(input)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
