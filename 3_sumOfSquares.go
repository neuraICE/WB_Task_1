package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

//Вывод суммы квадратов с помощью пакета sync
func sumOfSquares1(arr []int) {
	wg := new(sync.WaitGroup)
	mx := new(sync.Mutex)
	var result int

	for _, i := range arr {
		wg.Add(1)
		go func(j int, w *sync.WaitGroup, m *sync.Mutex) {
			m.Lock()
			defer m.Unlock()
			result += (j * j)
			w.Done()
		}(i, wg, mx)
	}
	wg.Wait()
	fmt.Println(result)
}

//Вывод суммы квадратов с помощью канала
func sumOfSquares2(arr []int) {
	ch := make(chan int)
	var result int

	for _, i := range arr {
		go func(j int, c chan int) {
			c <- j * j
		}(i, ch)
	}

	for range arr {
		result += <-ch
	}
	fmt.Printf("%d\t", result)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	sumOfSquares1(arr)
	fmt.Println()
	sumOfSquares2(arr)

}
