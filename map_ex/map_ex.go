package map_ex

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	m := map[string]int{}
	for _, v := range fields {
		m[v] += 1
	}

	return m
}

func run() {
	println("3. MAP:")
	wc.Test(WordCount)
}

func Run() {
	run()
}
