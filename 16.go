//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import "fmt"

//берем первое значение за отправную точку
//перебираем элементы
//делим на > и < отправной точки
//рукурсивно сортируем < и > элементы

func quckSort(input []int) []int {
	l := len(input)
	if l < 2 {
		return input
	}
	less := make([]int, 0)
	more := make([]int, 0)
	tochka := input[0]
	for _, v := range input[1:] {
		if v > tochka {
			more = append(more, v)
		} else {
			less = append(less, v)
		}
	}
	input = append(quckSort(less), tochka)
	input = append(input, quckSort(more)...)
	return input
}

func main() {
	array := []int{10, 4, 1, 6, 2, 5, 8, 3, 9, 7, 0}
	fmt.Println(quckSort(array))
}
