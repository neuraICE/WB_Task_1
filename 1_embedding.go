package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

type Human struct {
	Name
	Age
	Action
}

type Action struct {
	Work
	Rest
}

type Name string
type Age int
type Work string
type Rest string

func (a Action) say() {
	fmt.Printf("I work as a %s, and after work I like to %s.\n", a.Work, a.Rest)
}

func (a Action) write() {
	fmt.Println("I am writing a script\n")
}

func main() {
	human1 := Human{
		Name:   "Alex",
		Age:    34,
		Action: Action{"driver", "play to X-box and sleep"},
	}

	human2 := Human{
		Name:   "Selena",
		Age:    29,
		Action: Action{"stylist", "drink at the bar with friends"},
	}

	//Методы структуры Action используют обьекты родительской структуры Human:
	human1.say()
	human2.say()
	human2.write()

}
