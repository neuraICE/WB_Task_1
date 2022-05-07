package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func exchange(x, y *int) {
	*x, *y = *y, *x
}

func main() {

	var x = 1
	var y = 2

	fmt.Printf("x = %d\ny = %d\n", x, y)

	exchange(&x, &y)

	fmt.Printf("x = %d\ny = %d\n", x, y)
}
