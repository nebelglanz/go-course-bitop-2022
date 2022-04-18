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
	} //обработка ввода необходимого числа параметров
	log.Println("Выходной файл будет записан в output.json")
	log.Println("Делаю запрос")
	start := time.Now()                                                                                //начинаем отсчет времени
	resp, err := http.Get("https://bmstu.net/application/frontend/skin/one/assets/css/style.css?v=80") //получаем запрос
	if err != nil {
		log.Println("Не удалось совершить запрос", err)
		os.Exit(2) //обработка получения запроса
	} else {
		log.Println("Запрос сделан")
		log.Println("Запрос длился: ", time.Since(start))
	}
	var rawResp []byte
	rawResp, err = ioutil.ReadAll(resp.Body) //читаем запрос
	if err != nil {
		log.Println("Не удалось прочитать файл", err)
		os.Exit(2)
	} //обработка чтения запроса
	var sum, umn rune = 0, 1
	for _, value := range string(rawResp) {
		if string(value) >= "0" && string(value) <= "9" {
			//fmt.Print(value-48, " ")
			sum += value - 48
			umn *= value - 48
		}
	} //подсчет всех чисел из запроса
	result := fmt.Sprintf("Amount: %v\nProduct: %v", sum, umn) //запись в json файл
	err = ioutil.WriteFile(os.Args[1], []byte(result), 0755)
	if err != nil {
		log.Println("Не удалось записать в JSON-файл", err)
		os.Exit(2) //обработка записи в файл
	} else {
		log.Println("Скрипт выполнен")
	}
}
