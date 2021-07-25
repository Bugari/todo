package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `Add new task.
You can use tags #like #this`,
	Run: func(cmd *cobra.Command, args []string) {
		task := db.Task{Name: args[0]}
		handleAdd(&task)
		fmt.Printf("Created Task %s: %s\n", task.ID.String(), task.Name)
	},
}

func handleAdd(task *db.Task) (tx *gorm.DB) {
	return db.Conn.Create(&task)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
