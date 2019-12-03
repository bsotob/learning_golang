package main

import "fmt"

func main() {
	//suma(2, 4)
	//a := sumaRetorno(4, 4)
	//fmt.Println(a)
	//fmt.Println(sumaCombinada(2, 3, 4))
	fmt.Println(sumaCombinadaNombrada(2, 3, 4))
}
func suma(x int, y int) {
	fmt.Println(x + y)
}

func sumaRecortada(x, y int) {
	fmt.Println(x + y)
}
func sumaRetorno(x int, y int) int {
	return (x + y)
}
func sumaCombinada(x int, y int, z int) (int, int) {
	return x + y, x + z
}
func sumaCombinadaNombrada(x int, y int, z int) (l int, j int) {
	l = x + y
	j = x + z
	return

}
