//К каким негативным последствиям может привести данный фрагмент кода,
//и как это исправить? Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//v := createHugeString(1 << 10)
//justString = v[:100]
//}
//
//func main() {
//someFunc()
//}

package main

import "fmt"

var justString string

//Строки в GO это слайс байт
//Некоторые символы занимают не 1 байт (например кириллица или иероглифы)
//Поэтому нужно перевести строку в срез рун
//Пример сделан на генерации 1024 символов буквы Ж, каждая занимает 2 байта

//генерация строги
func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s += "Ж"
	}
	return
}

//вывод первых 100 символов
func someFunc() {
	a := createHugeString(1 << 10)
	b := []rune(a)
	//Разное кол-во символов в байтах и рунах. Поэтому переводим строку в руны

	fmt.Println(len(a))
	fmt.Println(len(b))

	justString = string(a[:100])
	fmt.Println(justString)
}
func main() {
	someFunc()
}
