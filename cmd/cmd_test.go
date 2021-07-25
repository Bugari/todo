package cmd_test

import (
	"fmt"
	"testing"
	. "todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands tests")
}

var _ = Describe("Commands", func() {
	BeforeEach(func() {
		db.OpenDb(":memory:?cache=shared")
	})
	Context("Handle add", func() {
		It("should add task properly to the db", func() {
			taskToAdd := db.Task{Name: "task"}

			var countExpected int64 = 1
			var countGot int64 = -1

			x := HandleAdd(&taskToAdd)

			Expect(x.RowsAffected).To(Equal(countExpected), "row is added")

			db.Conn.Model(&db.Task{}).Where("name = ?", taskToAdd.Name).Count(&countGot)
			Expect(countExpected).To(Equal(countGot), "task name exists in the db")
		})

		It("should add multiple tasks properly", func() {
			var tasks = []db.Task{{Seq: 1, Name: "task 1"}, {Seq: 2, Name: "task 2"}, {Seq: 3, Name: "task 3"}}
			var countAfter int64 = 0

			for _, t := range tasks {
				HandleAdd(&t)
			}

			db.Conn.Model(&db.Task{}).Where("1 = 1").Count(&countAfter)

			Expect(countAfter).To(Equal(int64(len(tasks))), "all taks should be in db")
		})
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
