package main

import (
	"fmt"
	"time"
)

func looper() {
	for i := 0; i < 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d \n", i)
	}
}

func main() {
	fmt.Println("start")

	go looper()
	fmt.Println("middle")
	go looper()

	var i int
	fmt.Scan(&i)
}
