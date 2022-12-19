//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

// разбиваем строку по пробелу на массив слов
// меняем их местами и затем объединяем обратно в строку через пробелы
func reverseWords(str string) string {
	s := strings.Split(str, " ")
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, " ")
}

func main() {
	str := "snow dog sun "
	fmt.Println(reverseWords(str))
}
