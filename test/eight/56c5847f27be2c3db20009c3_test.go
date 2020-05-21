package eight_test

import (
	. "codewars/eight"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subtract the Sum", func() {
	It("Basic tests", func() {
		Expect(SubtractSum(10)).To(Equal("apple"))
		Expect(SubtractSum(325)).To(Equal("apple"))
	})
})
