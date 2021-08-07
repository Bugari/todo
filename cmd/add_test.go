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
	Context("Priorities", func() {
		store := func(priority string) *db.Task {
			_, err := HandleAdd("task", priority, nil)
			Expect(err).To(BeNil())

			var retrievedTask db.Task
			db.Conn.Model(&db.Task{}).First(&retrievedTask)
			return &retrievedTask
		}
		It("should store high priority correctly", func() {
			task := store("high")
			Expect(task.GetPriority()).To(Equal("High"))
		})
		It("should store medium priority correctly", func() {
			task := store("med")
			Expect(task.GetPriority()).To(Equal("Medium"))
		})
		It("should store low priority correctly", func() {
			task := store("lo")
			Expect(task.GetPriority()).To(Equal("Low"))
		})
		It("should store default priority correctly", func() {
			task := store("")
			Expect(task.GetPriority()).To(Equal(""))
		})
	})
})
