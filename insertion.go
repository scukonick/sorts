package main

// SortInsertion sorts slice of int
// using insertion algorithm
func SortInsertion(a []int) {
	var i, j, buffer, countMove int
	for i = 1; i < len(a); i++ {
		buffer = a[i]
		countMove = 0
		for j = i - 1; j >= 0; j-- {
			if a[j] > buffer {
				countMove++
			} else {
				break
			}
		}
		if countMove > 0 {
			src := a[i-countMove : i]
			dst := a[i-countMove+1 : i+1]
			copy(dst, src)
			a[i-countMove] = buffer
		}
	}
}
