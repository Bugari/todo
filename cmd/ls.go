package cmd

import (
	"os"
	"time"
	"todo/db"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var renderDone bool = false

func ScopeDone(db *gorm.DB) *gorm.DB {
	if renderDone {
		return db.Where("done is not null")
	} else {
		return db.Where("done is null")
	}
}

const dayFormat = "2006-01-02"
const hourFormat = "15:04"
const zeroHour = "00:00"

func formatDue(task *db.Task) string {
	due := ""
	if task.Due == nil {
		return due
	}

	nowDate := time.Now().Format(dayFormat)
	formattedDate := task.Due.Format(dayFormat)
	formattedHour := task.Due.Format(hourFormat)
	if formattedDate != nowDate {
		due += formattedDate
	} else {
		due += "today"
	}
	if formattedHour != zeroHour {
		if due != "" {
			due += " "
		}
		due += formattedHour
	}
	return due
}
func printTasks(tasks *[]db.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Seq", "Due", "Priority", "Name"})

	rows := make([]table.Row, len(*tasks))
	for i, task := range *tasks {

		rows[i] = table.Row{task.Seq, formatDue(&task), task.GetPriority(), task.Name}
	}

	t.AppendRows(rows)
	t.SetStyle(table.StyleColoredDark)
	t.Render()
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list tasks",
	Long: `List tasks.
Listing updates ordering for following commands.
Currently lists only undone tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []db.Task
		db.Conn.Scopes(ScopeDone).Order("priority asc, created_at").Find(&tasks)
		if err := resetSeq(&tasks); err != nil {
			panic(err)
		}
		printTasks(&tasks)
	},
}

func resetSeq(tasks *[]db.Task) error {
	tx := db.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&db.Task{}).Where("1 = 1").Update("seq", nil).Error; err != nil {
		return err
	}

	for i := 0; i < len(*tasks); i++ {
		thisTask := &(*tasks)[i]
		thisTask.Seq = i + 1
		if err := tx.Save(&thisTask).Error; err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

func init() {
	lsCmd.Flags().BoolVar(&renderDone, "done", false, "display only finished tasks")
	rootCmd.AddCommand(lsCmd)
}
