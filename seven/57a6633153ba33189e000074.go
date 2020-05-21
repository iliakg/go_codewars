package seven

// Tuple Use the preloaded Tuple struct as return type
type Tuple struct {
	Char  rune
	Count int
}

// OrderedCount returns []Tuple
func OrderedCount(text string) []Tuple {
	result := []Tuple{}

	for _, char := range text {
		tuple := findTuple(result, char)
		if tuple == nil {
			result = append(result, Tuple{char, 1})
		} else {
			tuple.Count++
		}
	}
	return result
}

func findTuple(tuples []Tuple, char rune) (tup *Tuple) {
	for i, v := range tuples {
		if v.Char == char {
			tup = &tuples[i]
		}
	}
	return
}
