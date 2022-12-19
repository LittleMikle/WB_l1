package main

import "fmt"

//конвейер чисел
//принимает массив и канал в который будет запись из массива
func in(array []int, ch1 chan int) {
	defer close(ch1)
	for _, val := range array {
		ch1 <- val
	}
}

//принимает данные мз канала ch1 и записывает в ch2
func out(ch2 chan int, ch1 chan int) {
	defer close(ch2)
	for v := range ch1 {
		ch2 <- v * 2
	}

}

func main() {
	// каналы ch1 ch2 и массив из которого берем числа для записи
	ch1 := make(chan int)
	ch2 := make(chan int)
	array := []int{1, 2, 3, 4, 5}

	//запуск горутин
	go in(array, ch1)
	go out(ch2, ch1)

	//вывод из ch2
	for v := range ch2 {
		fmt.Println(v)
	}
}
