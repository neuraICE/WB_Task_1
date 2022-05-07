package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

// Функция removeByIndex1 удаляет элемент массива по индексу, сохраняя порядок
func removeByIndex1(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

// Функция removeByIndex2 удаляет элемент массива по индексу, не сохраняя порядок
func removeByIndex2(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	s := []string{"Bob", "Jim", "Mike", "john"}

	fmt.Println(removeByIndex2(s, 1))
	fmt.Println(removeByIndex1(s, 1))

}
