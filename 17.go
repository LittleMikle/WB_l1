//Реализовать бинарный поиск встроенными методами языка.
package main

import "fmt"

func SearchInIntSlice(haystack []int, needle int) (result bool, iterationCount int) {
	lowKey := 0                  //первый индекс
	highKey := len(haystack) - 1 //последний индекс
	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return //нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		//уменьшение списка рекурсивно
		iterationCount++
		mid := (lowKey + highKey) / 2 //середина
		if haystack[mid] == needle {
			result = true //значение найдено
			return
		}
		if haystack[mid] < needle {
			//если поиск больше середины - берем только с большими значениями увеличивая lowKey
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle {
			//если поиск меньше середины - берем только с большими значениями уменьшая highKey
			highKey = mid - 1
		}
	}
	return
}

func main() {
	result, count := SearchInIntSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6)
	fmt.Println(result, count)
}
