package eight_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eight Suite")
}
