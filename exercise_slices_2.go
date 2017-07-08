package main

import "golang.org/x/tour/pic"
```
結果：https://gyazo.com/58fa5e1f836e2298faf9f49613b58d5d
```
func Pic(dx, dy int) [][]uint8 {
	var shape [][]uint8
	var row = make([]uint8, dx)

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			row[j] = uint8(i * j)
		}
		shape = append(ret, row)
	}

	return shape
}

func main() {
	pic.Show(Pic)
}
