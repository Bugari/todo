package cmd

import (
	"fmt"
	"strings"
	"time"

	"todo/db"

	ft "github.com/bcampbell/fuzzytime"
	"github.com/spf13/cobra"
)

type AddArgs struct {
	Name        string
	PriorityRaw string
	NoPriority  bool
	DueDate     string
	NoDue       bool
}

var addArgs = AddArgs{}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `Add new task.
You can use tags #like #this`,
	Run: func(cmd *cobra.Command, args []string) {
		addArgs.Name = strings.Join(args, " ")
		task, err := HandleAdd(&addArgs)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created Task %s: %s\n", task.ID.String(), task.Name)
	},
}

func ParseDateArgs(primary string, secondary string) (*time.Time, error) {
	var dt ft.DateTime
	var err error

	if primary != "" {
		dt, _, err = ft.Extract(primary)
		if err != nil || dt.Empty() {
			return nil, err
		}
	} else {
		dt, _, err = ft.Extract(secondary)
		if err != nil || dt.Empty() {
			return nil, err
		}
	}

	formattedDate := fuzzytimeToDateTime(&dt)
	return formattedDate, nil
}

func fuzzytimeToDateTime(fuzzy *ft.DateTime) *time.Time {
	now := time.Now()
	year, month, day, hour, minute, second, nsec := now.Year(), now.Month(), 0, 0, 0, 0, 0
	if fuzzy.HasYear() {
		year = fuzzy.Year()
	}
	if fuzzy.HasMonth() {
		month = time.Month(fuzzy.Month())
	}
	if fuzzy.HasDay() {
		day = fuzzy.Day()
	}
	if fuzzy.HasHour() {
		hour = fuzzy.Hour()
	}
	if fuzzy.HasMinute() {
		minute = fuzzy.Minute()
	}
	if fuzzy.HasSecond() {
		second = fuzzy.Second()
	}

	time := time.Date(year, month, day, hour, minute, second, nsec, now.Location())
	return &time
}

func HandleAdd(args *AddArgs) (*db.Task, error) {
	task := db.Task{}
	applyArgsToTask(&task, args)
	result := db.Conn.Create(&task)
	return &task, result.Error
}

func applyArgsToTask(task *db.Task, args *AddArgs) {
	if args.Name != "" {
		task.Name = args.Name
	}
	if args.NoDue && task.Due != nil {
		task.Due = nil
	}
	date, err := ParseDateArgs(addArgs.DueDate, args.Name)
	if !args.NoDue && err == nil {
		task.Due = date
	}
	if args.NoPriority {
		task.SetPriority("")
	} else if args.PriorityRaw != "" {
		task.SetPriority(args.PriorityRaw)
	}
}

func attachAddParams(cmd *cobra.Command, params *AddArgs) {
	cmd.Flags().StringVarP(&params.PriorityRaw, "priority", "p", "", "Priority: high/medium/low h/m/l")
	cmd.Flags().BoolVarP(&params.NoPriority, "no-priority", "P", false, "Force no priority")
	cmd.Flags().StringVarP(&params.DueDate, "due", "d", "", "Due date")
	cmd.Flags().BoolVarP(&params.NoDue, "no-due", "D", false, "Force no due date")

}
func init() {
	attachAddParams(addCmd, &addArgs)
	rootCmd.AddCommand(addCmd)
}
