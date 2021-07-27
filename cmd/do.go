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

		task, err := HandleDone(taskNum)
		if task == nil {
			fmt.Printf("Task #%d not found\n", taskNum)
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("Task #%d marked as done: %s\n", task.Seq, task.Name)
	},
}

func HandleDone(taskNum int) (*db.Task, error) {
	var task db.Task
	result := db.Conn.First(&task, "seq = ?", taskNum)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	now := time.Now()
	task.Done = &now
	result = db.Conn.Save(&task)
	return &task, result.Error
}

func init() {
	rootCmd.AddCommand(doCmd)
}
