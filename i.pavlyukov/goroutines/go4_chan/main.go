package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
	ch chan []byte
)

func req(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ch <- body
}

func display(n int) {
	for {
		body := <-ch

		fmt.Println(n, string(body)[:10])

		wg.Done()
	}
}

func main() {
	ch = make(chan []byte)

	urls := []string{
		"https://yandex.ru",
		"https://google.com",
		"https://spatecon.ru",
		"https://vk.com",
		"https://wikipedia.org",
	}

	wg.Add(len(urls))
	for i := 0; i < len(urls); i++ {
		go req(urls[i])
	}

	fmt.Println(string(<-ch)[:10])
	fmt.Println(string(<-ch)[:10])
	fmt.Println(string(<-ch)[:10])
	fmt.Println(string(<-ch)[:10])
	fmt.Println(string(<-ch)[:10])
	fmt.Println(string(<-ch)[:10])

	fmt.Println("123")

	wg.Wait()
}
