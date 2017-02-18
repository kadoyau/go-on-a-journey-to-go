package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var old float64

	for i := 0; (old-z)*(old-z) > 0.00000001; i++ {
		old = z
		z = z - (z*z-x)/2*z
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2) - Sqrt(2))
}
