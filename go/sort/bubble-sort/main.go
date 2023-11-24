package main

import (
	"fmt"
)

func main() {
	tgt := []int{5, 4, 3, 7, 8, 1, 0}
	run(tgt)
	fmt.Println(tgt)
}

// 隣り合う要素を比較し、入れ替えながら整列する
func run(tgt []int) {
	for n := 0; n < len(tgt); n++ {
		var swap bool
		for j := 0; j < len(tgt)-n-1; j++ {
			if tgt[j] > tgt[j+1] {
				tgt[j], tgt[j+1] = tgt[j+1], tgt[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
