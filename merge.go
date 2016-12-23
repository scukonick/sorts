package main

func split(r []int) ([]int, []int) {
	l := len(r)
	sep := l / 2
	left := r[0:sep]
	right := r[sep:l]
	return left, right
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
OuterLoop:
	for i := 0; i < len(result); i++ {
		switch {
		case len(left) == 0:
			copy(result[i:len(result)], right)
			break OuterLoop

		case len(right) == 0:
			copy(result[i:len(result)], left)
			break OuterLoop

		default:
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:len(left)]
			} else {
				result[i] = right[0]
				right = right[1:len(right)]
			}
		}
	}
	return result
}

// SortMerge sortes input slice of integers
// using merge algorithm
func SortMerge(input []int) {
	left, right := split(input)
	if len(left) >= 2 {
		SortMerge(left)
	}
	if len(right) >= 2 {
		SortMerge(right)
	}
	result := merge(left, right)
	copy(input, result)

}
