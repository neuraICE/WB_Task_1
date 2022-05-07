package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

// Функция searchInSlice ищет значение в отсортированном слайсе разделяя его пополам
func searchInSlice(data []int, x int) (result bool, iterationsCount int) {
	sort.Ints(data)

	firstIndex := 0
	lastIndex := len(data) - 1

	if (data[firstIndex] > x) || (data[lastIndex] < x) {
		return
	}

	for firstIndex <= lastIndex {
		iterationsCount++

		midIndex := (firstIndex + lastIndex) / 2

		if data[midIndex] == x {
			result = true
			return
		}

		if data[midIndex] > x {
			lastIndex = midIndex - 1
		}

		if data[midIndex] < x {
			firstIndex = midIndex + 1
		}
	}
	return
}

func main() {

	arr1 := []int{32, 54, 2, 9, 5, 0, 44, 85, 66, 1, 29, 4}
	arr2 := []int{2, 9, 0, 43, 75, 83, 12, 38, 95, 101}

	fmt.Println(searchInSlice(arr1, 9))

	fmt.Println(searchInSlice(arr2, 4))
}
