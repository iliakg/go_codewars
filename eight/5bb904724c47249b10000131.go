package eight

import "strings"

// Points returns int
func Points(games []string) (points int) {
	for _, score := range games {
		splited := strings.Split(score, ":")
		points += point(&splited[0], &splited[1])
	}

	return
}

func point(a, b *string) int {
	if *a > *b {
		return 3
	} else if *a < *b {
		return 0
	}

	return 1
}
