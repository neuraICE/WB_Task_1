package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

//Вывод квадратов чисел с помощью канала
func valuesInSquares1(arr []int) {
	ch := make(chan int)

	for _, i := range arr {
		go func(j int, c chan int) {
			c <- j * j
		}(i, ch)
	}

	for range arr {
		fmt.Printf("%d\t", <-ch)
	}
}

//Вывод квадратов чисел с помощью WaitGroup
func valuesInSquares2(arr []int) {
	wg := new(sync.WaitGroup)

	for _, i := range arr {
		wg.Add(1)
		go func(j int, w *sync.WaitGroup) {
			fmt.Printf("%d\t", j*j)
			w.Done()
		}(i, wg)
	}
	wg.Wait()
}

func main() {

	arr := []int{2, 4, 6, 8, 10}

	valuesInSquares1(arr)
	fmt.Println()
	valuesInSquares2(arr)

}
