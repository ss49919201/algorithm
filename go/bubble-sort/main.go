package main

import (
	"fmt"
)

func main() {
	tgt := []int{5, 4, 3, 7, 8, 1, 0}
	run(tgt)
	fmt.Println(tgt)
}

func run(tgt []int) {
	if len(tgt) < 2 {
		return
	}
	compareAndSwap(tgt)
	run(tgt[:len(tgt)-1])
}

func compareAndSwap(tgt []int) {
	if len(tgt) < 2 {
		return
	}
	if tgt[0] > tgt[1] {
		tgt[0], tgt[1] = tgt[1], tgt[0]
	}
	compareAndSwap(tgt[1:])
}
