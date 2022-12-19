//Удалить i-ый элемент из слайса.
package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int
	fmt.Println("Введите номер элемента для удаления (от 0 до 10)")
	fmt.Scan(&i)
	fmt.Println("Слайс до изменения:")
	fmt.Println(array)
	//Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(array[i:], array[i+1:])
	//Удалить последний элемент (записать нулевое значение).
	array[len(array)-1] = 0
	//Усечь срез.
	array = array[:len(array)-1]
	fmt.Println("Слайс после изменения:")
	fmt.Println(array)
}
