package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	shape := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		shape[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			shape[i][j] = uint8(i)
		}
	}
	return shape
}

func main() {
	pic.Show(Pic)
}
