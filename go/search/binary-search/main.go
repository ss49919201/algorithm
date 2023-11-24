package main

import "fmt"

var arr []int
var tgt int

func main() {
	arr = []int{1, 3, 5, 7, 10, 13, 17}
	tgt = 9
	fmt.Println(run(0, len(arr)-1))

	tgt = 7
	fmt.Println(run(0, len(arr)-1))

	tgt = 3
	fmt.Println(run(0, len(arr)-1))
}

func run(low, high int) bool {
	if low > high {
		return false
	}

	mid := (high + low) / 2
	midv := arr[mid]

	switch {
	case midv == tgt:
		return true
	case midv < tgt:
		return run(mid+1, high)
	default:
		return run(low, mid-1)
	}
}
