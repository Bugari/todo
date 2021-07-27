package cmd_test

import (
	"testing"
	"todo/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands tests")
}

func pickOnlyNames(tasks []db.Task) []string {
	names := []string{}

	for _, task := range tasks {
		names = append(names, task.Name)
	}

	return names
}
