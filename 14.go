//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.
package main

import "fmt"

func getType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case chan interface{}:
		fmt.Println("chan interface{}:", v)
	default:
		fmt.Println("Unknown type:", v)

	}
}

func main() {
	channel := make(chan interface{})

	getType(0)
	getType("hello")
	getType(true)
	getType(channel)
	getType(0.99999)

}
