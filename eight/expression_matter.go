package eight

import "sort"

// ExpressionMatter returns int
func ExpressionMatter(a int, b int, c int) int {
	res := []int{
		a * b * c,
		a + b + c,
		(a + b) * c,
		a * (b + c),
		a + b*c,
		a*b + c,
	}

	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res[0]
}
