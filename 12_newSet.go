package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество
*/

// Функция newSet, итерируясь по входным данным, создает множество в map и возвращает его,
// где ключом являются сами данные а числовым значением - их исходное количество
func newSet(data []string) map[string]int {
	result := make(map[string]int)

	for _, i := range data {
		result[i]++
	}

	return result
}

func main() {

	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(newSet(arr))
}
