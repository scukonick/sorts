package main

import "testing"

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestSortInsertion(t *testing.T) {
	a := []int{5, 6, 1, 2, 3, 0, 4}
	SortInsertion(a)
	sorted := []int{0, 1, 2, 3, 4, 5, 6}
	if !testEq(a, sorted) {
		t.Errorf("Got: %v, expected: %v", a, sorted)
	}
}

func BenchmarkSortInsertion100(b *testing.B) {
	benchmarkSortFunc(b, 100, SortInsertion)
}

func BenchmarkSortInsertion1000(b *testing.B) {
	benchmarkSortFunc(b, 1000, SortInsertion)
}

func BenchmarkSortInsertion10000(b *testing.B) {
	benchmarkSortFunc(b, 10000, SortInsertion)
}
