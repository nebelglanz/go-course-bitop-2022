package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Введите название JSON-файла")
		os.Exit(2)
	}
	log.Println("Выходной файл будет записан в output.json")
	log.Println("Делаю запрос")
	start := time.Now()
	resp, err := http.Get("https://habr.com/kek/v2/articles/?company=vk&fl=ru&hl=ru&page=1")
	if err != nil {
		log.Println("Не удалось совершить запрос", err)
		os.Exit(2)
	} else {
		log.Println("Запрос сделан")
		log.Println("Запрос длился: ", time.Since(start))
	}
	var rawResp []byte
	rawResp, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Не удалось прочитать файл", err)
		os.Exit(2)
	}
	var sum, umn rune = 0, 1
	for _, value := range string(rawResp) {
		if value-48 >= 0 && value-48 <= 9 {
			//fmt.Print(value-48, " ") //вывести все цифры из запроса
			sum += value - 48
			umn *= value - 48
		}
	}
	result := []byte(fmt.Sprintf("Amount: %v\nProduct: %v", sum, umn))
	err = ioutil.WriteFile(os.Args[1], result, 0755)
	if err != nil {
		log.Println("Не удалось записать в JSON-файл", err)
		os.Exit(2)
	} else {
		log.Println("Скрипт выполнен")
	}
}
