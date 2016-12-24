package main

import "math"

func heapGetParent(i int) int {
	f := float64(i)
	parentF := math.Floor((f - 1) / 2)
	return int(parentF)
}

func heapGetLeftChild(i int) int {
	return 2*i + 1
}

func heapGetRightChild(i int) int {
	return 2*i + 2
}

func heapSwitch(input []int, pos int) {
	cValue := input[pos]
	p := heapGetParent(pos)
	pValue := input[p]
	if pValue < cValue {
		input[p], input[pos] = input[pos], input[p]
		r := heapGetRightChild(pos)
		l := heapGetLeftChild(pos)
		if r < len(input) {
			heapSwitch(input, r)
		}
		if l < len(input) {
			heapSwitch(input, l)
		}
	}
}

// SortHeap sorts input integer slice
// using Heap algorithm
func SortHeap(input []int) {
	if len(input) < 2 {
		return
	}
	// creating sorting tree
	for j := len(input) - 1; j >= 1; j-- {
		heapSwitch(input, j)
	}
	// moving element
	end := len(input) - 1
	input[end], input[0] = input[0], input[end]
	// repeat
	SortHeap(input[0 : len(input)-1])
}
