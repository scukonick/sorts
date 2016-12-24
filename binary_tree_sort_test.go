package main

import (
	"math/rand"
	"testing"
	"time"
)

type sortFunc func([]int)

func TestBinaryTreeSort(t *testing.T) {
	a := []int{5, 6, 1, 2, 3, 0, 4}
	SortBinaryTree(a)
	sorted := []int{0, 1, 2, 3, 4, 5, 6}
	if !testEq(a, sorted) {
		t.Errorf("Got: %v, expected: %v", a, sorted)
	}
}

func benchmarkSortFunc(b *testing.B, length int, f sortFunc) {
	a := make([]int, length)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl := make([]int, len(a))
		copy(sl, a[0:len(a)])
		f(sl)
	}
}

func BenchmarkSortBinaryTree100(b *testing.B) {
	benchmarkSortFunc(b, 100, SortBinaryTree)
}

func BenchmarkSortBinaryTree1000(b *testing.B) {
	benchmarkSortFunc(b, 1000, SortBinaryTree)
}

func BenchmarkSortBinaryTree10000(b *testing.B) {
	benchmarkSortFunc(b, 10000, SortBinaryTree)
}
