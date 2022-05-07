package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором.
*/

type Point struct {
	x int
	y int
}

func newPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// метод searchDistance рассчитывает дистанцию по формуле d = √((х А – х В) 2 + (у А – у В) 2)
func (p Point) searchDistance(p2 *Point) float64 {
	return math.Sqrt(float64((p.x-p2.x)*2 + (p.y-p2.y)*2))
}

func main() {
	p1, p2 := newPoint(4, 4), newPoint(2, 2)

	distance := p1.searchDistance(p2)

	fmt.Printf("%.2f\n", distance)
}
