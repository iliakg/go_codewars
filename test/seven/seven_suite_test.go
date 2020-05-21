package seven_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSeven(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Seven Suite")
}
