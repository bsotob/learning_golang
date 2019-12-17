package main

import (
	"context"
	"fmt"
)

//ctx, cancel := context.WithCancel(context.Background)
//context.Background genera un contexto vacío
//withCancel recibe un contexto base, que nos devuelve un contexto nuevo(ctx) y una funcion de cancelacion

func devuelveString(ctx context.Context) string {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	return "fin"
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go devuelveString(ctx)
	for i := 0; i < 20; i++ {
		fmt.Print("hola")
	}
	defer cancel() //cuando llega aqui, cancela el contexto.
}
