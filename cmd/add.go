package cmd

import (
	"fmt"
	"time"

	"todo/db"

	ft "github.com/bcampbell/fuzzytime"
	"github.com/spf13/cobra"
)

var priorityRaw = ""
var dueDate = ""
var ignoreDue = false

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `Add new task.
You can use tags #like #this`,
	Run: func(cmd *cobra.Command, args []string) {
		var due *time.Time
		if !ignoreDue {
			due, _ = ParseDateArgs(dueDate, args[0])
		}
		task, err := HandleAdd(args[0], priorityRaw, due)
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

func HandleAdd(name string, priority string, due *time.Time) (*db.Task, error) {
	task := db.Task{Name: name} //, Due:
	task.SetPriority(priority)
	task.Due = due
	result := db.Conn.Create(&task)
	return &task, result.Error
}

func init() {
	addCmd.Flags().StringVarP(&priorityRaw, "priority", "p", "", "Priority: high/medium/low h/m/l")
	addCmd.Flags().StringVarP(&dueDate, "due", "d", "", "Due date")
	addCmd.Flags().BoolVarP(&ignoreDue, "no-due", "D", false, "Ignore due dates in task description")
	rootCmd.AddCommand(addCmd)
}
