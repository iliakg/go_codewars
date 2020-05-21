package eight

import (
	"strconv"
	"strings"
)

var fruits = []string{
	"kiwi",
	"pear",
	"kiwi",
	"banana",
	"melon",
	"banana",
	"melon",
	"pineapple",
	"apple",
	"pineapple",
	"cucumber",
	"pineapple",
	"cucumber",
	"orange",
	"grape",
	"orange",
	"grape",
	"apple",
	"grape",
	"cherry",
	"pear",
	"cherry",
	"pear",
	"kiwi",
	"banana",
	"kiwi",
	"apple",
	"melon",
	"banana",
	"melon",
	"pineapple",
	"melon",
	"pineapple",
	"cucumber",
	"orange",
	"apple",
	"orange",
	"grape",
	"orange",
	"grape",
	"cherry",
	"pear",
	"cherry",
	"pear",
	"apple",
	"pear",
	"kiwi",
	"banana",
	"kiwi",
	"banana",
	"melon",
	"pineapple",
	"melon",
	"apple",
	"cucumber",
	"pineapple",
	"cucumber",
	"orange",
	"cucumber",
	"orange",
	"grape",
	"cherry",
	"apple",
	"cherry",
	"pear",
	"cherry",
	"pear",
	"kiwi",
	"pear",
	"kiwi",
	"banana",
	"apple",
	"banana",
	"melon",
	"pineapple",
	"melon",
	"pineapple",
	"cucumber",
	"pineapple",
	"cucumber",
	"apple",
	"grape",
	"orange",
	"grape",
	"cherry",
	"grape",
	"cherry",
	"pear",
	"cherry",
	"apple",
	"kiwi",
	"banana",
	"kiwi",
	"banana",
	"melon",
	"banana",
	"melon",
	"pineapple",
	"apple",
	"pineapple",
}

// SubtractSum returns string
func SubtractSum(n int) string {
	for ok := true; ok; ok = n > 100 {
		sum := 0
		for _, st := range strings.Split(strconv.Itoa(n), "") {
			i, _ := strconv.Atoi(st)
			sum += i
		}
		n -= sum
	}

	return fruits[n-1]
}