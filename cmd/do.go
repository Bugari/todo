package cmd

import (
	"fmt"
	"strconv"
	"time"

	"todo/db"

	"github.com/spf13/cobra"
)

// doCmd represents the add command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as done",
	Long:  `Mark task as done`,
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

		fmt.Printf("Task #%d marked as done: %s\n", task.Seq, task.Name)
	},
}

func HandleDone(task *db.Task) {
	now := time.Now()
	task.Done = &now
	db.Conn.Save(&task)
}

func init() {
	rootCmd.AddCommand(doCmd)
}
