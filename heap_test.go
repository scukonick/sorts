package main

import "testing"

func TestSortHeap(t *testing.T) {
	a := []int{5, 6, 1, 2, 3, 0, 4, 50, 124, 4}
	SortHeap(a)
	sorted := []int{0, 1, 2, 3, 4, 4, 5, 6, 50, 124}
	if !testEq(a, sorted) {
		t.Errorf("Got: %v, expected: %v", a, sorted)
	}
}

func BenchmarkSortHeap100(b *testing.B) {
	benchmarkSortFunc(b, 100, SortHeap)
}

func BenchmarkSortHeap1000(b *testing.B) {
	benchmarkSortFunc(b, 1000, SortHeap)
}

func BenchmarkSortHeap10000(b *testing.B) {
	benchmarkSortFunc(b, 10000, SortHeap)
}
