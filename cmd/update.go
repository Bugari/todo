package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"todo/db"

	"github.com/spf13/cobra"
)

type UpdateArgs struct {
	AddArgs
	TargetSeqNum int
}

var updateArgs = UpdateArgs{}

// addCmd represents the add command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update task",
	Long:  `update task`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		updateArgs.TargetSeqNum, err = strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Sprintf("Incorrect task number: \"%s\"\n", args[0]))
		}
		if len(args) > 1 {
			updateArgs.Name = strings.Join(args[1:], " ")
		}
		task, err := HandleUpdate(&updateArgs)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created Task %s: %s\n", task.ID.String(), task.Name)
	},
}

func HandleUpdate(args *UpdateArgs) (*db.Task, error) {
	var task db.Task
	result := db.Conn.Model(&db.Task{}).First(&task, "seq = ?", args.TargetSeqNum)
	if result.RowsAffected == 0 {
		panic(fmt.Sprintf("Task not found: \"%d\"\n", args.TargetSeqNum))
	}
	applyArgsToTask(&task, &args.AddArgs)
	result = db.Conn.Save(&task)

	return &task, result.Error
}

func attachUpdateParams(cmd *cobra.Command, params *UpdateArgs) {
}

func init() {
	attachAddParams(updateCmd, &updateArgs.AddArgs)
	attachUpdateParams(updateCmd, &updateArgs)
	rootCmd.AddCommand(updateCmd)
}
