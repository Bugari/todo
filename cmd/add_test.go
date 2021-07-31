package cmd_test

import (
	"time"
	. "todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
	BeforeEach(func() {
		db.OpenDb(":memory:")
	})
	Context("Handle add", func() {
		It("should add task properly to the db", func() {
			taskToAdd := db.Task{Name: "task"}

			var countExpected int64 = 1
			var countGot int64 = -1

			_, err := HandleAdd("task", "", nil)

			Expect(err).To(BeNil(), "error while adding task")

			db.Conn.Model(&db.Task{}).Where("name = ?", taskToAdd.Name).Count(&countGot)
			Expect(countExpected).To(Equal(countGot), "task name does not exist in the db")
		})

		It("should add multiple tasks properly", func() {
			var tasks = []string{"task 1", "task 2", "task 3"}
			var countAfter int64 = 0

			for _, t := range tasks {
				HandleAdd(t, "", nil)
			}

			db.Conn.Model(&db.Task{}).Where("1 = 1").Count(&countAfter)

			Expect(countAfter).To(Equal(int64(len(tasks))), "all taks should be in db")
		})
	})
	Context("parseDate", func() {
		It("should parse secondary properly", func() {
			now := time.Now()
			expected := now.Format("02 Jan 06")
			date, _ := ParseDateArgs("", now.Format("02 Jan 06"))
			Expect(date).ToNot(BeNil())
			Expect(date.Format("02 Jan 06")).To(Equal(expected))
		})
		It("should parse primary properly", func() {
			now := time.Now()
			expected := now.Format("02 Jan 06")
			date, _ := ParseDateArgs(now.Format("02 Jan 06"), "")
			Expect(date).ToNot(BeNil())
			Expect(date.Format("02 Jan 06")).To(Equal(expected))
		})
		It("should pass nil on empty dates", func() {
			date, _ := ParseDateArgs("", "")
			Expect(date).To(BeNil())
		})
	})

})
