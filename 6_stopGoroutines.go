package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

const COUNT = 5

//Остановка горутины с помощью пакета sync
func stopGoroutine1() {
	wg := new(sync.WaitGroup)
	wg.Add(COUNT)
	go func(w *sync.WaitGroup) {
		for i := 0; i < COUNT; i++ {
			fmt.Println("Print line:", i)
			w.Done()
		}
	}(wg)
	fmt.Println("Goroutine stopped")
	wg.Wait()
	fmt.Println("Goroutine completed")
}

//Остановка горутины с помощью пакета time
func stopGoroutine2() {
	fmt.Println("Goroutine start")
	go func() {
		for i := 0; i < COUNT; i++ {
			fmt.Println("Print line:", i)
		}
	}()
	time.Sleep(time.Second * COUNT)
	fmt.Println("Goroutine completed")
}

//Остановка горутины с помощью пакета context
func stopGoroutine3() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err := operation1(ctx); err != nil {
			cancel()
		}
	}()
	operation2(ctx)
}

func operation1(ctx context.Context) error {
	time.Sleep(time.Second)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(time.Second * COUNT):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

//Остановка горутины с помощью канала
func stopGoroutine4() {
	ch := make(chan string)

	go func(c chan string) {
		fmt.Println(<-c)
		time.Sleep(2 * time.Second)
		c <- "Goroutine 2 stopped"
	}(ch)

	ch <- "Goroutine 1 stopped"

	fmt.Println(<-ch)
}

//Остановка горутины с помощью канала типа struct{}
func stopGoroutine5() {
	arr := []int{1, 2, 6, 7, 8, 5, 4, 0, 22, 65, 9}

	done := make(chan struct{})
	defer close(done)

	for i := range squa(done, gener(done, arr)) {
		fmt.Println(i)
	}
}

// Функция gener записывает данные в канал из массива до получения сигнала done
func gener(done <-chan struct{}, arr []int) <-chan int {
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

// Функция squa читает данные из канала и, удваивая, записывает их в другой канал до получения сигнала done
func squa(done <-chan struct{}, in <-chan int) <-chan int {
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
