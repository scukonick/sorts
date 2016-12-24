package main

import (
	"math/rand"
	"testing"
	"time"
)

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

func benchmarkSortInsertion(b *testing.B, length int) {
	a := make([]int, length)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl := make([]int, len(a))
		copy(sl, a[0:len(a)])
		SortInsertion(sl)
	}
}

func BenchmarkSortInsertion100(b *testing.B) {
	benchmarkSortInsertion(b, 100)
}

func BenchmarkSortInsertion1000(b *testing.B) {
	benchmarkSortInsertion(b, 1000)
}

func BenchmarkSortInsertion10000(b *testing.B) {
	benchmarkSortInsertion(b, 10000)
}
