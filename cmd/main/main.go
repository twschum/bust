package main

import (
	"bust/pkg/dice"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	total := 0
	for roll := range dice.RollGen(3, 6) {
		fmt.Println(roll)
		total += 1
	}
	fmt.Printf("total rolls: %d", total)
}
