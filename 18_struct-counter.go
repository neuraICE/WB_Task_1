package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Count struct {
	sync.Mutex
	int
}

// Метод inc блокирует доступ другим горутинам к полю структуры с помощью mutex во время инкрементации
func (c *Count) inc() {
	c.Lock()
	c.int++
	c.Unlock()
}

var count Count

func main() {
	wg := new(sync.WaitGroup)

	amount := 10000
	wg.Add(amount)

	for i := 0; i < amount; i++ {
		go func(i int, w *sync.WaitGroup) {
			defer w.Done()
			count.inc()
		}(i, wg)
	}

	wg.Wait()
	fmt.Println(count.int)
}
