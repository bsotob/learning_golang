package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//tambien existe switch-case con evalucaciones
	ahora := time.Now().Weekday()
	fmt.Println(ahora)
	switch time.Sunday {
	case ahora + 0:
		fmt.Println("Hoy")
	case ahora + 1:
		fmt.Println("Mañana")
	case ahora + 2:
		fmt.Println("Pasado mañana")
	default:
		fmt.Println("Quedan mas de 3 dias")
	}

	//fin switch case con evaluaciones

	system := runtime.GOOS
	switch system {
	case "darwin":
		fmt.Println("macos")
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("no soportado")
	}

	/* hacemos lo mismo de abajo con swicht
	system := runtime.GOOS
	if system == "darwin" {
		fmt.Println("macos")
	}
	if system == "linux" {
		fmt.Println("linux")
	}
	if system == "windows" {
		fmt.Println("windows")
	}
	*/
}
