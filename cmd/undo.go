package cmd

import (
	"fmt"
	"strconv"

	"todo/db"

	"github.com/spf13/cobra"
)

// undoCmd represents the add command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Mark task as undone",
	Long:  `Mark task as undone`,
	Run: func(cmd *cobra.Command, args []string) {
		taskNum, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Sprintf("Incorrect task number: \"%s\"\n", args[0]))
		}

		task, err := HandleUndo(taskNum)
		if task == nil {
			fmt.Printf("Task #%d not found\n", taskNum)
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("Task #%d marked as undone: %s\n", task.Seq, task.Name)
	},
}

func HandleUndo(taskNum int) (*db.Task, error) {
	var task db.Task
	result := db.Conn.First(&task, "seq = ?", taskNum)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	task.Done = nil
	result = db.Conn.Save(&task)
	return &task, result.Error
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
