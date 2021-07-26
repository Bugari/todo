package cmd_test

import (
	"fmt"
	. "todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleDone", func() {
	var tasks = []db.Task{{Seq: 1, Name: "a"}, {Seq: 2, Name: "b"}, {Seq: 3, Name: "c"}}
	BeforeEach(func() {
		db.OpenDb(":memory:")
		db.Conn.Create(&tasks)
	})
	Context("Marking task as done", func() {
		It("should have tasks in db", func() {
			var countExpected int64 = int64(len(tasks))
			var countGot int64 = -1

			db.Conn.Model(&db.Task{}).Where("1 = 1").Count(&countGot)

			Expect(countExpected).To(Equal(countGot), "DB tasks count do not match")
		})

		It("should not be done by default", func() {
			var task db.Task
			db.Conn.Model(&db.Task{}).First(&task)

			Expect(task.Done).To(BeZero(), "Done field is nil")
		})
		It("should have all tasks not done yet", func() {
			arr := []db.Task{}
			result := db.Conn.Model(&db.Task{}).Find(&arr, "done is null")
			Expect(result.RowsAffected).To(Equal(int64(len(tasks))), "DB should have all tasks not marked as done")
		})

		It("should mark first task as done", func() {
			var task db.Task
			db.Conn.Model(&db.Task{}).First(&task)

			changedTask, _ := HandleDone(1)

			Expect(changedTask.Done).NotTo(BeZero(), "Done field is not nil")
		})

		for i, task := range tasks {
			It(fmt.Sprintf("should not return task #%d\n", i), func() {
				arr := []db.Task{}
				db.Conn.Model(&db.Task{}).Find(&arr, "done is null")

				HandleDone(task.Seq)

				db.Conn.Model(&db.Task{}).Find(&arr, "done is null")
				Expect(len(arr)).To(Equal(len(tasks)-1), "Done field is not nil")

				Expect(arr).NotTo(WithTransform(pickOnlyNames, ContainElement(task.Name)), fmt.Sprintf("task should not be returned anymore %d\n", i))
			})

		}

	})
})
