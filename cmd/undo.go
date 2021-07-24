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

		var task db.Task

		result := db.Conn.First(&task, "seq = ?", taskNum)
		if result.RowsAffected == 0 {
			panic(fmt.Sprintf("Unknown task with number: \"%s\"\n", args[0]))
		}

		task.Done = nil
		db.Conn.Save(&task)

		fmt.Printf("Task #%d marked as undone: %s\n", task.Seq, task.Name)
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
