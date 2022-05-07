package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

// Функция gen записывает данные в канал из массива до получения сигнала done
func gen(done <-chan struct{}, arr []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, i := range arr {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()

	return out
}

// Функция sq читает данные из канала и, удваивая, записывает их в другой канал до получения сигнала done
func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			select {
			case out <- i * 2:
			case <-done:
				return
			}
		}
		close(out)
	}()

	return out
}

// С помощью канала done завершаем работу остальных горутин, когда данные становятся не актуальны
func main() {

	arr := []int{1, 2, 6, 7, 8, 5, 4, 0, 22}

	done := make(chan struct{})
	defer close(done)

	for i := range sq(done, gen(done, arr)) {
		fmt.Println(i)
	}

}
