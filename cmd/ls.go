package cmd

import (
	"os"
	"todo/db"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func printTasks(tasks *[]db.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Seq", "Name"})

	rows := make([]table.Row, len(*tasks))
	for i, task := range *tasks {
		rows[i] = table.Row{task.ID, task.Seq, task.Name}
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
		db.Conn.Find(&tasks, "done is null")
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
		thisTask := (*tasks)[i]
		thisTask.Seq = i + 1
		if err := tx.Save(&thisTask).Error; err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
