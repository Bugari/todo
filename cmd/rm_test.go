package cmd

import (
	"os"
	"testing"
	"todo/db"
)

func TestMain(m *testing.M) {
	db.OpenDb(":memory:")
	code := m.Run()
	os.Exit(code)
}

func Test_handleRm(t *testing.T) {
	t.Run("remove task properly", func(t *testing.T) {
		taskToRemove := db.Task{Seq: 1, Name: "task"}
		db.Conn.Create(&taskToRemove)

		handleRm(1)

		var countExpected int64 = 0
		var countGot int64 = -1

		db.Conn.Model(&db.Task{}).Where("seq = 1").Count(&countGot)
		if countExpected != countGot {
			t.Errorf("Expected task was not removed.k Found %d matching tasks.", countGot)
		}
	})
	t.Run("remove just one correct task", func(t *testing.T) {
		var tasks = []db.Task{{Seq: 1, Name: "task 1"}, {Seq: 2, Name: "task 2"}, {Seq: 3, Name: "task 3"}}
		db.Conn.Create(&tasks)

		handleRm(2)

		var remainingTasks []db.Task
		db.Conn.Find(&remainingTasks).Order("seq")
		if len(remainingTasks) != 2 {
			t.Errorf("Expected 2 tasks, got %d", len(remainingTasks))
		}
		if remainingTasks[0].Seq != 1 || remainingTasks[1].Seq != 3 {
			t.Error("Wrong task was removed.")
		}

	})
}
