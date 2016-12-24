package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	var a [50]int
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % 20
	}
	sl := a[0:len(a)]
	log.Printf("Before: %v", sl)
	SortHeap(sl)
	log.Printf("After: %v", sl)
}
