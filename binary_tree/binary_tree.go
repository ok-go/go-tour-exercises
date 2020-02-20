package binary_tree

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		i, ok1 := <-ch1
		j, ok2 := <-ch2

		if i != j || ok1 != ok2 {
			return false
		}

		if !ok1 && !ok2 {
			return true
		}
	}
}

func testWalk() {
	print("Test Walk(): ")
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	j := 1
	for i := range ch {
		if i != j {
			println("failed")
			return
		}
		j++
	}
	println("success")
}

func testSame() {
	r1 := Same(tree.New(1), tree.New(1))
	r2 := !Same(tree.New(1), tree.New(2))

	print("Test Same(): ")
	if r1 && r2 {
		println("success")
	} else {
		println("failed")
		println("\ttree.New(1) == tree.New(1):", r1)
		println("\ttree.New(1) != tree.New(2):", r2)
	}
}

func run() {
	println("10. BINARY TREE:")
	testWalk()
	testSame()
}

func Run() {
	run()
}
