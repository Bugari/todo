package cmd_test

import (
	. "todo/cmd"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
	BeforeEach(func() {
		db.OpenDb(":memory:")
	})
	Context("Handle update", func() {
		It("should update name properly", func() {
			tmpTask := db.Task{Seq: 1, Name: "task"}
			db.Conn.Create(&tmpTask)

			HandleUpdate(&UpdateArgs{AddArgs: AddArgs{Name: "new name"}, TargetSeqNum: 1})

			var newTask db.Task
			db.Conn.Model(&db.Task{}).First(&newTask, "seq = ?", 1)

			Expect(newTask.Name).To(Equal("new name"))
		})
	})
})
