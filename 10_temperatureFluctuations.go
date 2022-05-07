package main

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

// Функция grouping, используя оператор switch распределяет данные в соответствующие группы
func grouping(arr []float64) map[int][]float64 {
	result := make(map[int][]float64)
	for _, i := range arr {
		switch {
		case i >= -30 && i < -20:
			result[-30] = append(result[-30], i)
		case i >= -20 && i < -10:
			result[-20] = append(result[-20], i)
		case i >= -10 && i < 0:
			result[-10] = append(result[-10], i)
		case i >= 0 && i < 10:
			result[0] = append(result[0], i)
		case i >= 10 && i < 20:
			result[10] = append(result[10], i)
		case i >= 20 && i < 30:
			result[20] = append(result[20], i)
		case i >= 30 && i < 40:
			result[30] = append(result[30], i)
		}
	}
	return result
}

func main() {

	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(grouping(temp))
}
