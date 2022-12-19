//Разработать программу, которая перемножает, делит, складывает, вычитает две
//числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(1000000000)
	y := big.NewInt(200000000)
	z := big.NewInt(0)

	fmt.Println("Результат умножения:", z.Mul(x, y))
	fmt.Println("Результат деления:", z.Div(x, y))
	fmt.Println("Результат сложения:", z.Add(x, y))
	fmt.Println("Результат вычитания:", z.Sub(x, y))
}
