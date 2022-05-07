package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

// Функция sleep первым аргументом получает длительность,
// а вторым название периода времени ('s' - секунды, 'm' - миллисекунды)
// и останавливает горутину на заданный период
func sleep(value int, t string) {
	switch t {
	case "m":
		time.Sleep(time.Millisecond * time.Duration(value))
	case "s":
		time.Sleep(time.Second * time.Duration(value))
	}
}

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("I am sleeping %d second\n", i)
		sleep(i, "s")
	}

	for i := 100; i <= 500; i += 100 {
		fmt.Printf("I am sleeping %d millisecond\n", i)
		sleep(i, "m")
	}
}
