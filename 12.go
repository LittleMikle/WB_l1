//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для
//нее собственное множество.
package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool) //итоговое множество

	for _, word := range words {
		if _, ok := set[word]; !ok { //если слова нет во множестве, добавляем
			set[word] = true
		}
	}

	for k, _ := range set {
		fmt.Println(k)
	}
}

