package main

import (
	"fmt"
	"strings"
)

// используем мапу для проверки уникальных символов
func checkUnique(str string) bool {
	str = strings.ToLower(str)
	runeds := []rune(str)
	unique := make(map[rune]struct{})

	for _, char := range runeds {
		if _, ok := unique[char]; ok {
			return false
		}
		unique[char] = struct{}{}
	}
	return true
}

func main() {
	var testString string
	fmt.Println("Введите строку для проверки")
	fmt.Scan(&testString)
	fmt.Println(checkUnique(testString))

}
