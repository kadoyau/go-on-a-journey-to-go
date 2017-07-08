package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	counter := 0
	n_0 := 0
	n_1 := 0
	sum := 0
	return func() int {
		if counter == 0 {
			sum = 0
		} else if counter == 1 {
			sum = 1
			n_1 = sum
		} else {
			sum = n_0 + n_1
			n_0 = n_1
			n_1 = sum
		}
		counter += 1
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
