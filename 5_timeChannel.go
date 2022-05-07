package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения
в канал, а с другой стороны канала — читать. По истечению N секунд
программа должна завершаться.
*/

const N = 4

func someProgram() {
	ch := make(chan int)

	ctx, _ := context.WithTimeout(context.Background(), N*time.Second)

	go func(c chan int) {
		for {
			select {
			case <-time.After(500 * time.Millisecond):
				c <- rand.Intn(10)
			case <-ctx.Done():
				close(c)
				return
			}
		}
	}(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
