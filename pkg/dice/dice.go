package dice

// RollGen is a generator-style function for sequentially generating all the possible
// dice rolls for a given number of dice with the same number of sides
func RollGen(dice int, sides int) chan []int {
	out := make(chan []int)
	go func() {
		//atomic.AddUint64(&sparked, 1)
		defer close(out)
		// generate this iterations dice
		//for d := range dieGen(sides) {
		for d := 1; d < sides+1; d++ {
			if dice > 1 {
				for remainder := range RollGen(dice-1, sides) {
					roll := append([]int{d}, remainder...)
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
		//atomic.AddUint64(&sparked, 1)
		defer close(out)
		for i := 1; i < sides+1; i++ {
			out <- i
		}
	}()
	return out
}
