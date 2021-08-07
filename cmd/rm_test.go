package cmd_test

import (
	"fmt"
	. "todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleRm", func() {
	BeforeEach(func() {
		db.OpenDb(":memory:?cache=shared")
	})
	Context("Handle rm", func() {
		It("remove task properly", func() {
			taskToRemove := db.Task{Seq: 1, Name: "task"}
			db.Conn.Create(&taskToRemove)

			HandleRm(1)

			var countExpected int64 = 0
			var countGot int64 = -1

			db.Conn.Model(&db.Task{}).Where("seq = 1").Count(&countGot)
			Expect(countGot).To(Equal(countExpected), fmt.Sprintf("Expected task was not removed. Found %d matching tasks.", countGot))
		})

		It("remove just one correct task", func() {
			var tasks = []db.Task{{Seq: 1, Name: "task 1"}, {Seq: 2, Name: "task 2"}, {Seq: 3, Name: "task 3"}}
			db.Conn.Create(&tasks)

			HandleRm(2)

			var remainingTasks []db.Task
			db.Conn.Find(&remainingTasks).Order("seq")
			Expect(remainingTasks).To(HaveLen(2), fmt.Sprintf("Expected 2 tasks, got %d", len(remainingTasks)))
			Expect(remainingTasks[0].Seq).To(Equal(1), "Wrong task was removed")
			Expect(remainingTasks[1].Seq).To(Equal(3), "Wrong task was removed")
		})
	})
})
