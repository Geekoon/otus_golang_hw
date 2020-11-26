package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ct := time.Now()
	et, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal("Something wrong with ntp: ", err)
	}

	fmt.Println("current time:", ct.Round(time.Nanosecond))
	fmt.Println("exact time:", et.Round(time.Nanosecond))
}
