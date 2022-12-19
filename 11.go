//Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mn1 := make([]int, 10)
	mn2 := make([]int, 10)

	var cross []int

	//заполняем наборы
	for i, _ := range mn1 {
		mn1[i] = rand.Intn(5)
	}
	for i, _ := range mn2 {
		mn2[i] = rand.Intn(5)
	}

	//поиск пересечения
	for _, mn1 := range mn1 {
		for _, mn2 := range mn2 {
			if mn1 == mn2 {
				cross = append(cross, mn1)
				break
			}
		}
	}
	fmt.Println("Set1:", mn1)
	fmt.Println("Set2:", mn2)
	fmt.Println("Set cross:", cross)
}
