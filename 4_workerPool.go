package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

// Структура для взаимодействия горутин
type control struct {
	wg   sync.WaitGroup
	sig  chan os.Signal
	data chan string
}

// Конструктор для создания объекта структуры
func newControl() *control {
	return &control{
		sig:  make(chan os.Signal, 1),
		data: make(chan string),
	}
}

// worker выводит данные из канала с небольшим ожиданием, до получения сигнала через контекст
func (c *control) worker(id int, ctx context.Context) {
	c.wg.Add(1)
	defer c.wg.Done()
	for {
		select {
		case i := <-c.data:
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Worker %d says: %s\n", id, i)
		case <-ctx.Done():
			fmt.Printf("Worker %d completed\n", id)
			return
		}
	}
}

// sendInChannel отправляет данные в канал генерируя их с помощью randomString(), до получения сигнала через контекст
func (c *control) sendInChannel(ctx context.Context) {
	defer close(c.data)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Data sending completed!\n")
			return
		default:
			time.Sleep(500 * time.Millisecond)
			c.data <- randomString()
		}
	}
}

// randomString генерирует данные типа string размера 10 символов
func randomString() string {
	bytes := make([]byte, 10)
	for i := range bytes {
		bytes[i] = byte(rand.Intn(26) + 65)
	}
	return string(bytes)
}

// waitingStopCommand ожидает сигнала "^control + C",
// после чего сигнализирует вызывающей горутине
func (c *control) waitingStopCommand() bool {
	sigs := <-c.sig
	fmt.Println(sigs)
	return true
}

func main() {

	// poolSize позволяет задавать количество воркеров при старте программы,
	// устанавливая значение флага "-pools" в терминале (по умолчанию количество равно 7)
	poolSize := flag.Int("pools", 7, "Количество воркеров")
	flag.Parse()

	pool := newControl()

	signal.Notify(pool.sig, syscall.SIGINT, syscall.SIGTERM)

	var ctx, cancel = context.WithCancel(context.Background())

	// запускаем воркеров
	for i := 1; i <= *poolSize; i++ {
		go pool.worker(i, ctx)
	}

	go pool.sendInChannel(ctx)

	go func() {
		if ok := pool.waitingStopCommand(); ok {
			cancel() // после получения ответа от waitingStopCommand завершаем остальные горутины через context
		}
	}()

	// ожидаем завершения всех горутин
	pool.wg.Wait()
	fmt.Println("Function main completed")
}
