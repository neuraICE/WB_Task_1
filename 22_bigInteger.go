package main

import (
	"fmt"
	"log"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

// Функция operation принимает два аргумента типа *big.Int и тип операции в string и возвращает результать
func operation(a, b *big.Int, o string) *big.Int {
	res := new(big.Int)
	switch o {
	case "+":
		res.Add(a, b)
	case "-":
		res.Sub(a, b)
	case "*":
		res.Mul(a, b)
	case "/":
		res.Div(a, b)
	case "%":
		res.Rem(a, b)
	default:
		log.Println("Операция не найдена!")
		return nil
	}
	return res
}

func main() {

	a, ok := new(big.Int).SetString("400000000000000000000", 10)
	if !ok {
		log.Println("создание переменной 'a' не удалось")
	}
	b, ok := new(big.Int).SetString("200000000000000000000", 10)
	if !ok {
		log.Println("создание переменной 'b' не удалось")
	}

	operations := []string{"+", "-", "*", "/", "%"}

	for _, i := range operations {
		fmt.Println(operation(a, b, i))
	}

}
