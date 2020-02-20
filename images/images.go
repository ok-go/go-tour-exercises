package images

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	ux, uy := uint8(x), uint8(y)
	return color.RGBA{ux, uy, 255, 255}
}

func run() {
	print("9. IMAGES: ")
	m := Image{}
	pic.ShowImage(m)
}

func Run() {
	run()
}
