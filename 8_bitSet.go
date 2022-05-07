package main

import (
	"errors"
	"fmt"
	"log"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// Функция bitSet принимает три аргумента: значение в десятичном формате типа int64, позицию изменяемого бита и новое значение бита.
// Далее возвращает измененное число в двоичном формате ввиде строки
func bitSet(value int64, b int, v int) (string, error) {
	if b < 1 {
		return "", errors.New("range exceeded")
	}
	if v == 1 {
		i := value | (1 << (b - 1))
		return fmt.Sprintf("%b", i), nil
	} else if v == 0 {
		i := value &^ (1 << (b - 1))
		return fmt.Sprintf("%b", i), nil
	} else {
		return "", errors.New("invalid value! Expected '1' or '0'")
	}
}

func main() {

	var x int64 = -9203072036800775807

	fmt.Printf("%b\n", x)

	y, err := bitSet(x, 62, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(y)
}
