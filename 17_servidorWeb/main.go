package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hola", func(w http.ResponseWriter, resp *http.Request) {
		fmt.Printf("%v , %v, %v", resp.Method, resp.Host, resp.Header)
		fmt.Fprintln(w, "Hola mundo")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error creando el servidor")
		fmt.Println(err)
	}
}
