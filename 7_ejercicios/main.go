package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//leemos un string usando bufio
	fmt.Println("añada su nombre: ")
	fmt.Println(" * * * * * * * *")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre := scanner.Text()
	if len(nombre) > 0 {
		fmt.Println("tu nombre es: ", scanner.Text())
	} else {
		fmt.Println("no hay ningun string")
	}

}
