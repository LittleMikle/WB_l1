//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//квадратов(22+32+42….) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//для каждого числа вызывается горутина, считает квадрат и складывает с atomic переменной
func Solution() {
	var array = []uint64{2, 4, 6, 8, 10}
	var result uint64
	var wg sync.WaitGroup
	for _, num := range array {
		wg.Add(1)
		go func(num uint64) {
			defer wg.Done()
			atomic.AddUint64(&result, num*num)
		}(num)
	}

	wg.Wait()
	fmt.Println(result)
}

func main() {
	Solution()
}
