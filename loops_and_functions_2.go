package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	diff := float64(10)
	z := float64(1)
	old := z

	for diff > 0.00000001 {
		z = z - (z*z-x)/2*z
		diff = (old - z) * (old - z)
		old = z
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2) - Sqrt(2))
}
