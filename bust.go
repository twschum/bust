package main

import (
	"fmt"
	"sort"
	"sync/atomic"
	//"golang.org/x/exp/slices"
)

var sparked uint64

func main() {
	fmt.Println("vim-go")
	total := 0
	for roll := range rollGen(3, 4) {
		//roll = slices.Sort(roll)
		fmt.Println(roll)
		total += 1
	}
	fmt.Printf("total rolls: %d\ntotal goroutines: %v", total, sparked)
}

// rollGen is a generator-style function for sequentially generating all the possible
// dice rolls for a given number of dice with the same number of sides
func rollGen(dice int, sides int) chan []int {
	out := make(chan []int)
	go func() {
		atomic.AddUint64(&sparked, 1)
		defer close(out)
		// generate this iterations dice
		//for d := range dieGen(sides) {
		for d := 1; d < sides+1; d++ {
			if dice > 1 {
				for remainder := range rollGen(dice-1, sides) {
					roll := append([]int{d}, remainder...)
					sort.Slice(roll, func(i, j int) bool {
						return roll[i] < roll[j]
					})
					out <- roll
				}
			} else {
				out <- []int{d}
			}
		}
	}()
	return out
}

func dieGen(sides int) chan int {
	out := make(chan int)
	go func() {
		atomic.AddUint64(&sparked, 1)
		defer close(out)
		for i := 1; i < sides+1; i++ {
			out <- i
		}
	}()
	return out
}
