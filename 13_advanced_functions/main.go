package main

import (
	"fmt"
	"math"
)

type FuncionesCalculadora func(float64, float64) float64

func computar(funcion func(int, int) int, a int, b int) int {
	return funcion(a, b)
}
func computar2(funcion FuncionesCalculadora, a float64, b float64) float64 {
	return funcion(a, b)
}

func main() {

	suma := func(a, b int) int {
		return a + b
	}
	resta := func(a, b int) int {
		return a - b
	}
	fmt.Println(suma(1, 2))
	fmt.Println(computar(suma, 1, 2))
	fmt.Println(computar(resta, 1, 2))
	fmt.Println(computar2(math.Pow, 2, 3))
}
