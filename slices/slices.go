package slices

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for y := range r {
		r[y] = make([]uint8, dx)
		for x := range r[y] {
			r[y][x] = uint8(x + y)
		}
	}

	return r
}

func run() {
	print("2. SLICES: ")
	pic.Show(Pic)
}

func Run() {
	run()
}
