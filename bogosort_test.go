package main

import "testing"

func TestSortBogosort(t *testing.T) {
	a := []int{5, 2}
	SortBogo(a)
	sorted := []int{2, 5}

	if !testEq(a, sorted) {
		t.Errorf("Got: %v, expected: %v", a, sorted)
	}
}
