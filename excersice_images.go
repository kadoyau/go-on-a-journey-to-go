package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct {
	sizeX int
	sizeY int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.sizeX, i.sizeY)
}

func (i Image) At(x, y int) color.Color {
	// Calculate y*255 before devide by sizeY because sizeY > 255 then 255/i.sizeY = 0
	return color.RGBA{0, uint8(y * 255 / i.sizeY), uint8(x * 255 / i.sizeX), 255}

}

func main() {
	m := Image{500, 255} // 500 x 255
	pic.ShowImage(m)
	// result: https://gyazo.com/685734a2c52d2b8d73ceded964eab68b
}
