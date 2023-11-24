package main

import "fmt"

func main() {
	fmt.Println(sort([]int{1, 5, 2, 20, 4, 3, 1}))
}

func sort(tgt []int) []int {
	// 空の場合は、そのまま返す
	if len(tgt) == 0 {
		return tgt
	}

	// 中間あたりを基準値にする
	x := len(tgt) / 2

	// 二つに分割
	var l, r []int
	for i, v := range tgt {
		// 基準値はスルー
		if i == x {
			continue
		}

		// 基準値以下とそれ以外に分ける
		if v <= tgt[x] {
			l = append(l, v)
		} else {
			r = append(r, v)
		}
	}

	// それぞれを再起的にソート
	l = sort(l)
	r = sort(r)

	// 基準値を基準値以下の末尾に追加
	// lは基準値以下なので、基準値をソート対象にする必要ない
	// つまり挿入はソート後でOK
	l = append(l, tgt[x])

	return append(l, r...)
}
