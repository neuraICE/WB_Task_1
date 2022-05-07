package main

import (
	"math/rand"
	"strings"
	"time"
)

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/
// В приведенном выше примере переменная justString ссылается на исходный массив/строку,
// сборщик мусора не может освободить массив, несколько полезных байтов хранят все содержимое в памяти.
// Эту проблему можно устранить, используя функцию append

var justString string

// Функция someFunc создает временный массив рун и добавляет туда часть исходной строки с помощью append,
// после чего добавляет данные в переменную justString
func someFunc() {
	v := createHugeString(1 << 10)

	var tempString []rune

	tempString = append(tempString, []rune(v[:100])...)

	justString = string(tempString)
}

// Функция createHugeString генерирует строку с заданным размером
func createHugeString(j int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var hugeString strings.Builder
	for i := 0; i < j; i++ {
		hugeString.WriteRune(chars[rand.Intn(len(chars))])
	}
	return hugeString.String()
}

func main() {
	someFunc()
}
