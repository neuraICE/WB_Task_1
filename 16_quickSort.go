package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

// С помощью финкции quickSort рекурсивно делим массив на более мелкие отрезки
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	split := partition(arr)

	quickSort(arr[:split])
	quickSort(arr[split:])
}

// С помощью финкции partition двигаемся к середине массива, меняя значения местами,
// которые находятся не на своем месте и возвращает индекс из середины
func partition(arr []int) int {
	pivot := arr[len(arr)/2]

	left := 0
	right := len(arr) - 1

	for {
		for ; arr[left] < pivot; left++ {
		}

		for ; arr[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}

func main() {

	arr := []int{67, 4, 5, 3, 56, 32, 0, 9, 8}

	quickSort(arr)

	fmt.Println(arr)
}
