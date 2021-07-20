package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list tasks",
	Long: `List all tasks
You will be told more
when time comes.`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []db.Task
		db.Conn.Find(&tasks)
		for _, task := range tasks {
			fmt.Printf("%d: %s\n", task.ID, task.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
