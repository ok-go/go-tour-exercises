package loop_and_func

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	p := .1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(p-z) < 10e-9 {
			break
		}
		p = z
	}

	return z
}

func cmp(x float64) {
	fmt.Println("Compare for", x)
	fmt.Println("My: ", Sqrt(x))
	fmt.Println("Std:", math.Sqrt(x))
}

func run() {
	println("1. LOOP_AND_FUNC:")
	cmp(1)
	cmp(2)
	cmp(3)
}

func Run() {
	run()
}
