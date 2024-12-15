package cmd

import (
	"strconv"

	"github.com/rajatnai/task/db"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove un-wanted task",
	Run: func(cmd *cobra.Command, args []string) {
		var input []int
		for _, v := range args {
			i, err := strconv.Atoi(v)
			if err == nil {
				input = append(input, i)
			}
		}
		db.DeleteTask(input)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
