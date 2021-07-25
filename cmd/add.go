package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
)

var priorityRaw = ""

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `Add new task.
You can use tags #like #this`,
	Run: func(cmd *cobra.Command, args []string) {
		task, err := HandleAdd(args[0], priorityRaw)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created Task %s: %s\n", task.ID.String(), task.Name)
	},
}

func HandleAdd(name string, priority string) (*db.Task, error) {
	task := db.Task{Name: name}
	task.SetPriority(priority)
	result := db.Conn.Create(&task)
	return &task, result.Error
}

func init() {
	addCmd.Flags().StringVarP(&priorityRaw, "priority", "p", "", "Priority: high/medium/low h/m/l")
	rootCmd.AddCommand(addCmd)
}
