//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской
//структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Sex     string
	Age     uint
	Height  uint
	Weight  float32
}

type Action struct {
	Human
}

func (h Human) GetName() string {
	return h.Name
}

func (h Human) GetSurname() string {
	return h.Surname
}

func (h Human) GetSex() string {
	return h.Sex
}

func (h Human) GetAge() uint {
	return h.Age
}

func (h Human) GetHeight() uint {
	return h.Height
}

func (h Human) GetWeight() float32 {
	return h.Weight
}

func main() {
	human := Human{"Mikle", "Melushev", "M", 23, 183, 67.89}
	action := Action{Human: human}
	fmt.Printf(" Name is %s \n Surname is %s \n Sex is %s \n Age is %d \n Height is %d \n Weight is %.2f \n", action.GetName(), action.GetSurname(), action.GetSex(), action.GetAge(), action.GetHeight(), action.GetWeight())
}

