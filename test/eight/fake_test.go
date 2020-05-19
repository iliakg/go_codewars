package eight_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Fake Test", func() {
	It("fake true", func() { Expect(true).To(Equal(true)) })
})
