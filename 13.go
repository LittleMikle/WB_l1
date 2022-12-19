//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	first := 1
	second := 2

	first, second = second, first
	fmt.Println(first, second)
}
