package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// counter количество работающих горутин
	counter int
	m       sync.Mutex
)

func foo(n int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("[%d]: chto-to %d\n", n, i)
	}

	m.Lock()
	counter--
	m.Unlock()
}

func main() {
	counter = 10
	for i := 0; i < counter; i++ {
		go foo(i)
	}

	for {
		if counter == 0 {
			break
		}
		time.Sleep(time.Second)
	}
}
