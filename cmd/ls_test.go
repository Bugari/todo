package cmd_test

import (
	"todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleLs", func() {
	BeforeEach(func() {
		db.OpenDb(":memory:?cache=shared")
	})
	Describe("resetSeq", func() {
		It("should have no sequence numbers at the start", func() {
			var tasks = []db.Task{{Name: "task 1"}, {Name: "task 2"}, {Name: "task 3"}}
			db.Conn.Create(&tasks)

			var recoveredTask db.Task
			db.Conn.First(&recoveredTask)

			Expect(recoveredTask.Seq).To(Equal(0))
		})
		It("should set sequence numbers", func() {
			var tasks = []db.Task{{Name: "task 1"}, {Name: "task 2"}, {Name: "task 3"}}
			db.Conn.Create(&tasks)
			cmd.ResetSeq(&tasks)

			var recoveredTasks []db.Task
			db.Conn.Find(&recoveredTasks).Order("seq asc")

			for i := 0; i < 3; i++ {
				Expect(recoveredTasks[i].Seq).To(Equal(i + 1))
			}
		})

	})
})
