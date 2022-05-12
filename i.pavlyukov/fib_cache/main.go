package main

import (
	"fmt"
	"time"
)

// 0 1 1 2 3 5 8 13 21	...
// 0 1 2 3 4 5 6 7 8	... 10000

var m map[int]int

func fib(i int) int {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	v, ok := m[i]
	if !ok {
		v = fib(i-1) + fib(i-2)
		m[i] = v
	}

	return v
}

func main() {
	started := time.Now()
	defer func() {
		fmt.Println(time.Since(started))
	}()

	m = make(map[int]int, 1000)

	fmt.Println(fib(1000))
}
