//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.
package main

import "fmt"

func stringReverse(str string) string {
	runeds := []rune(str)
	for i, j := 0, len(runeds)-1; i < len(runeds)/2; i, j = i+1, j-1 {
		runeds[i], runeds[j] = runeds[j], runeds[i]
	}
	return string(runeds)
}
func main() {
	var s string
	fmt.Println("Input your string:")
	fmt.Scan(&s)
	fmt.Println(stringReverse(s))
}
