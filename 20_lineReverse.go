package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

// Функция reverseL переворачивает строку, создавая массив строк с помощью пакета strings
func reverseL(s string) string {
	result := strings.Fields(s)

	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return strings.Join(result, " ")
}

func main() {

	line1 := "snow dog sun"
	line2 := "sea cat door"

	fmt.Println(reverseL(line1))
	fmt.Println(reverseL(line2))

}
