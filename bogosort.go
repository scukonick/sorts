package main

import "math/rand"

func check(sl []int) bool {
	for i := 1; i < len(sl); i++ {
		if sl[i-1] > sl[i] {
			return false
		}
	}
	return true
}

// SortBogo sort slice using bogosort algorithm
func SortBogo(input []int) {
	if len(input) <= 1 {
		return
	}

	shuffle := func(sl []int) {
		for i := range sl {
			j := rand.Intn(i + 1)
			sl[i], sl[j] = sl[j], sl[i]
		}
	}

	for {
		if check(input) {
			return
		}
		shuffle(input)
	}
}
