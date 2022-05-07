package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

// Функция definition использует пакет reflect для выявления типа полученного значения
func definition(data interface{}) {
	t := reflect.TypeOf(data)
	fmt.Println(t)
}

func main() {

	var (
		i int
		s string
		b bool
		c chan string
	)

	definition(i)
	definition(s)
	definition(b)
	definition(c)
}
