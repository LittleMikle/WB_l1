//Написать программу, которая конкурентно рассчитает значение квадратов
//чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
	"time"
)

var array = []int{2, 4, 6, 8, 10}

//решение 1
//создается буферизированный канал
//в цикле для каждого числа создается горутина и пишет в канал квадраты
//ожидание завершения через  wg.Wait,  закрывается канал и чтение из него
func firstSol() {
	result := make(chan int, len(array))
	var wg sync.WaitGroup
	for _, num := range array {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			result <- num * num
			fmt.Println(<-result)

		}(num)
	}
}

//решение 2
//горутина в цикле для каждого числа с выводом квадрата
//time.Sleep для ожидания завершения отработки всех горутин
func secondSol() {
	for _, num := range array {
		go func(num int) {
			fmt.Println(num * num)
		}(num)
	}
	time.Sleep(time.Second * 1)

}

//решение 3
//создается небуферизированный канал
//в цикле для каждого числа запускается горутина и записывает в канал квадрат
//чтение из канала len(array)
func thirdSol() {
	result := make(chan int)
	for i, _ := range array {
		go func(i int) {
			result <- array[i] * array[i]
		}(i)
	}
	for i := 0; i < len(array); i++ {
		fmt.Println(<-result)
	}
	close(result)
}

func main() {
	firstSol()
	secondSol()
	thirdSol()
}
