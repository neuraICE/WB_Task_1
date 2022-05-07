package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна
быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

// Функция characterCheck проверяет уникальность символов,
// добавляя их в качестве ключа в map и инкрементируя значение при итерациии
func characterCheck(s string) bool {
	m := make(map[rune]int)
	runes := []rune(strings.ToLower(s))

	for _, i := range runes {
		m[i]++
		if m[i] > 1 {
			return false
		}
	}
	return true
}

func main() {
	s := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, i := range s {
		fmt.Println(characterCheck(i))
	}
}
