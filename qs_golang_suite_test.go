package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQsGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "QsGolang Suite")
}
