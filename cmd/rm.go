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
		HandleRm(taskNum)
	},
}

func HandleRm(taskNum int) {
	var task db.Task
	result := db.Conn.First(&task, "seq = ?", taskNum)
	if result.RowsAffected != 1 {
		fmt.Printf("Task #%d not found\n", taskNum)
		return
	}

	if err := db.Conn.Delete(&task).Error; err != nil {
		panic(err)
	}

	// db.Conn.Save(&task)
	fmt.Printf("Task #%d removed: %s\n", task.Seq, task.Name)
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
