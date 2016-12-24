package main

import "testing"

func TestSortMerge(t *testing.T) {
	a := []int{5, 6, 1, 2, 3, 0, 4}
	SortMerge(a)
	sorted := []int{0, 1, 2, 3, 4, 5, 6}
	if !testEq(a, sorted) {
		t.Errorf("Got: %v, expected: %v", a, sorted)
	}
}

func TestSortMergeConcurrent(t *testing.T) {
	a := []int{5, 6, 1, 2, 3, 0, 4}
	SortMergeConcurrent(a)
	sorted := []int{0, 1, 2, 3, 4, 5, 6}
	if !testEq(a, sorted) {
		t.Errorf("Got: %v, expected: %v", a, sorted)
	}
}

func BenchmarkSortMerge100(b *testing.B) {
	benchmarkSortFunc(b, 100, SortMerge)
}

func BenchmarkSortMerge1000(b *testing.B) {
	benchmarkSortFunc(b, 1000, SortMerge)
}

func BenchmarkSortMerge10000(b *testing.B) {
	benchmarkSortFunc(b, 10000, SortMerge)
}

func BenchmarkSortMergeConcurrent100(b *testing.B) {
	benchmarkSortFunc(b, 100, SortMergeConcurrent)
}

func BenchmarkSortMergeConcurrent1000(b *testing.B) {
	benchmarkSortFunc(b, 1000, SortMergeConcurrent)
}

func BenchmarkSortMergeConcurrent10000(b *testing.B) {
	benchmarkSortFunc(b, 10000, SortMergeConcurrent)
}
