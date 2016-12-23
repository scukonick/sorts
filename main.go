package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	var a [5]int
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % 10
	}
	sl := a[0:len(a)]
	log.Printf("Before: %v", sl)
	SortMerge(sl)
	log.Printf("After: %v", sl)
}