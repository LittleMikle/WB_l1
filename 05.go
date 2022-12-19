//Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func Solution() {
	var t int
	fmt.Println("Set timer in seconds:")
	fmt.Scan(&t)

	var ch = make(chan int)                             //создание канала для чтение и записи
	timer := time.After(time.Duration(t) * time.Second) //таймер для остановки
	var wg sync.WaitGroup
	wg.Add(1)
	//запись в ch
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timer:
				log.Printf("Writing to ch stopped after %d seconds", t)
				close(ch)
				return
			default:
				randNum := rand.Intn(100)
				ch <- randNum
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Add(1)
	//чтение из сh
	go func() {
		defer wg.Done()
		for range ch {
			if val, ok := <-ch; ok {
				fmt.Printf("Writed %d\n", val)
			} else {
				fmt.Println("Channel closed !!!!!")
			}
		}
	}()

	wg.Wait()
}

func main() {
	Solution()
}
