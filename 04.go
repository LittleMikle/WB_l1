//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор
//из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		select {
		case <-ctx.Done():
			return
		default:
			if value, ok := <-ch; ok {
				fmt.Printf("worker %d received %d\n", id, value)
			} else {
				fmt.Println("Channel closed")
			}
		}
	}
}

func Solution() {
	var wg sync.WaitGroup
	var workersNum int
	ch := make(chan int)                         //главный канал
	signalChannel := make(chan os.Signal, 1)     //канал для SIGINT
	signal.Notify(signalChannel, syscall.SIGINT) //запись в канал для SIGINT если он пришел
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Set num of workers:")
	fmt.Scan(&workersNum)
	log.Println("Started")

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	//если был ввод CTRL + C происходит отмена контекста и закрытие каналов
	for {
		select {
		case <-signalChannel:
			cancel()
			close(ch)
			close(signalChannel)
			wg.Wait()
			log.Println("Program stopped")
			return
		default:
			ch <- rand.Intn(100)

			time.Sleep(time.Second)
		}
	}
}

func main() {
	Solution()
}
