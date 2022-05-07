package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

// В функции intersection пересечение реализуется посредством map
func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, e := range a {
		counter[e]++
	}

	for _, e := range b {
		if _, ok := counter[e]; ok {
			result = append(result, e)
		}
		counter[e]++
	}

	return result
}

func main() {

	a := []int{2, 3, 4, 44, 7, 8, 99, 0}
	b := []int{99, 6, 44, 2, 13, 7, 56, 88}

	for _, i := range intersection(a, b) {
		fmt.Println(i)
	}
}
