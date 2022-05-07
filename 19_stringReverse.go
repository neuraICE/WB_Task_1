package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

// Функция reverseS переворачивает строку, конвертируя ее в тип rune
func reverseS(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

func main() {

	s1 := "главрыба"
	s2 := "4/][.09=-1§±"

	fmt.Println(reverseS(s1))
	fmt.Println(reverseS(s2))
}
