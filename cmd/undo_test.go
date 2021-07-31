package cmd_test

import (
	"time"
	. "todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleDone", func() {
	doneTime := time.Now()
	sourceTask := db.Task{Seq: 1, Name: "a", Done: &doneTime}
	BeforeEach(func() {
		db.OpenDb(":memory:")
		db.Conn.Create(&sourceTask)
	})
	Context("Marking task as done", func() {
		It("should be done beforehand", func() {
			var task db.Task
			db.Conn.First(&task, "Done is not null")
			Expect(task).ToNot(BeNil(), "Task should be done before other tests")
		})

		It("should have tasks in db", func() {
			HandleUndo(1)

			var task db.Task
			db.Conn.First(&task)

			Expect(task.Done).To(BeNil(), "Task should be not done")
		})

	})
})
