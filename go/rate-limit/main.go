package main

import "time"

var logQ []time.Time

func main() {
	var exec bool
	var curWindow int
	if len(logQ) > 1 {
		if time.Since(logQ[len(logQ)-1]) <= 60 {
			curWindow++

			if time.Since(logQ[len(logQ)-2]) <= 60 {
				curWindow++
			}
		}
	}

	if curWindow < 2 {
		exec = true
		logQ = logQ[len(logQ)-curWindow+1:]
	}

	// ログの追記
	logQ = append(logQ, time.Now())

	// 実行
	if exec {
	}
}
