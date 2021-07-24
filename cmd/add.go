package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `Add new task.
You can use tags #like #this`,
	Run: func(cmd *cobra.Command, args []string) {
		task := db.Task{Name: args[0]}
		db.Conn.Create(&task)
		fmt.Printf("Created Task %d: %s", task.ID, task.Name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
