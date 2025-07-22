package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {

	var c float64 = float64(a) // Convertendo int para float64
	fmt.Println("Valor de c:", c)
	fmt.Printf("Valor de b: %s\n", b)
}
