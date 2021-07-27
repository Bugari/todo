package cmd

import (
	"fmt"
	"strconv"

	"todo/db"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove task",
	Long:  `Remove task from list`,
	Run: func(cmd *cobra.Command, args []string) {
		taskNum, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Sprintf("Incorrect task number: \"%s\"\n", args[0]))
		}
		task, err := HandleRm(taskNum)
		if task == nil {
			fmt.Printf("Task #%d not found\n", taskNum)
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("Task #%d removed: %s\n", task.Seq, task.Name)
	},
}

func HandleRm(taskNum int) (*db.Task, error) {
	var task db.Task
	result := db.Conn.First(&task, "seq = ?", taskNum)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	result = db.Conn.Delete(&task)
	return &task, result.Error
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
