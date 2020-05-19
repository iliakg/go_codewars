package eight_test

import (
	. "codewars/eight"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExpressionMatter", func() {
	It("ExpressionMatter(2, 1, 2)", func() { Expect(ExpressionMatter(2, 1, 2)).To(Equal(6)) })
	It("ExpressionMatter(2, 1, 1)", func() { Expect(ExpressionMatter(2, 1, 1)).To(Equal(4)) })
	It("ExpressionMatter(1, 1, 1)", func() { Expect(ExpressionMatter(1, 1, 1)).To(Equal(3)) })
	It("ExpressionMatter(1, 2, 3)", func() { Expect(ExpressionMatter(1, 2, 3)).To(Equal(9)) })
	It("ExpressionMatter(1, 3, 1)", func() { Expect(ExpressionMatter(1, 3, 1)).To(Equal(5)) })
	It("ExpressionMatter(2, 2, 2)", func() { Expect(ExpressionMatter(2, 2, 2)).To(Equal(8)) })
	It("ExpressionMatter(5, 1, 3)", func() { Expect(ExpressionMatter(5, 1, 3)).To(Equal(20)) })
	It("ExpressionMatter(3, 5, 7)", func() { Expect(ExpressionMatter(3, 5, 7)).To(Equal(105)) })
	It("ExpressionMatter(5, 6, 1)", func() { Expect(ExpressionMatter(5, 6, 1)).To(Equal(35)) })
	It("ExpressionMatter(1, 6, 1)", func() { Expect(ExpressionMatter(1, 6, 1)).To(Equal(8)) })
	It("ExpressionMatter(2, 6, 1)", func() { Expect(ExpressionMatter(2, 6, 1)).To(Equal(14)) })
	It("ExpressionMatter(6, 7, 1)", func() { Expect(ExpressionMatter(6, 7, 1)).To(Equal(48)) })
	It("ExpressionMatter(2, 10, 3)", func() { Expect(ExpressionMatter(2, 10, 3)).To(Equal(60)) })
	It("ExpressionMatter(1, 8, 3)", func() { Expect(ExpressionMatter(1, 8, 3)).To(Equal(27)) })
	It("ExpressionMatter(9, 7, 2)", func() { Expect(ExpressionMatter(9, 7, 2)).To(Equal(126)) })
	It("ExpressionMatter(1, 1, 10)", func() { Expect(ExpressionMatter(1, 1, 10)).To(Equal(20)) })
	It("ExpressionMatter(9, 1, 1)", func() { Expect(ExpressionMatter(9, 1, 1)).To(Equal(18)) })
	It("ExpressionMatter(10, 5, 6)", func() { Expect(ExpressionMatter(10, 5, 6)).To(Equal(300)) })
	It("ExpressionMatter(1, 10, 1)", func() { Expect(ExpressionMatter(1, 10, 1)).To(Equal(12)) })
})
