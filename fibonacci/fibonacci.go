package fibonacci

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pp := -1
	p := 1
	return func() int {
		curr := p + pp

		pp = p
		if p == 0 {
			p = 1
		} else {
			p = curr
		}

		return curr
	}
}

func run() {
	print("4. FIBONACCI: ")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	println()
}

func Run() {
	run()
}
