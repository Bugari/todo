package cmd_test

import (
	"fmt"
	. "todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleDone", func() {
	prios := []string{"High", "Medium", "Low"}
	tasksNames := []string{"a", "b", "c"}

	BeforeEach(func() {
		db.OpenDb(":memory:")
		for idx, task := range tasksNames {
			HandleAdd(task, prios[idx])
		}
	})
	Context("Marking task as done", func() {
		It("should have tasks in db", func() {
			var countExpected int64 = int64(len(tasksNames))
			var countGot int64 = -1

			db.Conn.Model(&db.Task{}).Where("1 = 1").Count(&countGot)

			Expect(countExpected).To(Equal(countGot), "DB tasks count do not match")
		})

		It("should have all tasks not done yet", func() {
			arr := []db.Task{}
			result := db.Conn.Model(&db.Task{}).Find(&arr, "done is null")
			Expect(result.RowsAffected).To(Equal(int64(len(tasksNames))), "DB should have all tasks not marked as done")
		})

		It("should mark first task as done", func() {
			arr := []db.Task{}
			db.Conn.Model(&db.Task{}).Find(&arr, "done is null")

			Expect(arr[0].Done).To(BeZero(), "Done field is nil")

			HandleDone(&arr[0])

			Expect(arr[0].Done).NotTo(BeZero(), "Done field is not nil")
		})

		for i, name := range tasksNames {
			It(fmt.Sprintf("should not return task #%d\n", i), func() {
				arr := []db.Task{}
				db.Conn.Model(&db.Task{}).Find(&arr, "done is null")

				Expect(arr[i].Done).To(BeZero(), "Done field is nil")

				HandleDone(&arr[i])

				db.Conn.Model(&db.Task{}).Find(&arr, "done is null")
				Expect(len(arr)).To(Equal(len(tasksNames)-1), "Done field is not nil")

				Expect(arr).NotTo(WithTransform(pickOnlyNames, ContainElement(name)), fmt.Sprintf("task should not be returned anymore %d\n", i))
			})

		}

	})
})
