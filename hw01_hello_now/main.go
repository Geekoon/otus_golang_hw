package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Узнаем текущее время на хосте
	ct := time.Now()

	// Узнаем точное время с использованием библиотеки ntp
	et, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	// Обрабатываем ошибку обращения к удаленному серверу
	if err != nil {
		log.Fatal("Something wrong with ntp: ", err)
	}

	// Выводим результат с округлением времени, потраченного на вычисление (монотонное время)
	fmt.Println("current time:", ct.Round(time.Nanosecond))
	fmt.Println("exact time:", et.Round(time.Nanosecond))
}
