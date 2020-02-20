package errors

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x
	p := .1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(p-z) < 10e-9 {
			break
		}
		p = z
	}

	return z, nil
}

func run() {
	println("6. ERRORS:")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Run() {
	run()
}
