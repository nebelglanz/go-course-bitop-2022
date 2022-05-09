package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
)

func firstBytes(url string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(url, string(body)[:10])
}

func main() {
	urls := []string{
		"https://yandex.ru",
		"https://google.com",
		"https://spatecon.ru",
		"https://vk.com",
		"https://wikipedia.org",
	}

	wg.Add(len(urls))
	for i := 0; i < len(urls); i++ {
		go firstBytes(urls[i])
	}

	wg.Wait()
}
