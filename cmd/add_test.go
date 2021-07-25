package cmd

import (
	"testing"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "handleAdd")
}

var _ = Describe("handleAdd", func() {
	BeforeSuite(func() {
		db.OpenDb(":memory:?cache=shared")
	})
	BeforeEach(func() {
		db.Conn.Unscoped().Delete(&db.Task{}, "1 = 1")
	})
	AfterSuite(func() {
		db.Conn.Unscoped().Delete(&db.Task{}, "1 = 1")
	})
	It("should add task properly to the db", func() {
		taskToAdd := db.Task{Name: "task"}

		var countExpected int64 = 1
		var countGot int64 = -1

		x := handleAdd(&taskToAdd)

		Expect(x.RowsAffected).To(Equal(countExpected), "row is added")

		db.Conn.Model(&db.Task{}).Where("name = ?", taskToAdd.Name).Count(&countGot)
		Expect(countExpected).To(Equal(countGot), "task name exists in the db")
	})

	It("should add multiple tasks properly", func() {
		var tasks = []db.Task{{Seq: 1, Name: "task 1"}, {Seq: 2, Name: "task 2"}, {Seq: 3, Name: "task 3"}}
		var countAfter int64 = 0

		for _, t := range tasks {
			handleAdd(&t)
		}

		db.Conn.Model(&db.Task{}).Where("1 = 1").Count(&countAfter)

		Expect(countAfter).To(Equal(int64(len(tasks))), "all taks should be in db")
	})
})
