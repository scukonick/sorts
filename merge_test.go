package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestSortMerge(t *testing.T) {
	a := []int{5, 6, 1, 2, 3, 0, 4}
	SortMerge(a)
	sorted := []int{0, 1, 2, 3, 4, 5, 6}
	if !testEq(a, sorted) {
		t.Errorf("Got: %v, expected: %v", a, sorted)
	}
}

func benchmarkSortMerge(b *testing.B, length int) {
	a := make([]int, length)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl := make([]int, len(a))
		copy(sl, a[0:len(a)])
		SortMerge(sl)
	}
}

func BenchmarkSortMerge100(b *testing.B) {
	benchmarkSortMerge(b, 100)
}

func BenchmarkSortMerge1000(b *testing.B) {
	benchmarkSortMerge(b, 1000)
}

func BenchmarkSortMerge10000(b *testing.B) {
	benchmarkSortMerge(b, 10000)
}
