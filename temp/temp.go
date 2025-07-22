package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {

	var tempF = ebulicaoF
	tempC := (tempF - 32) * 5 / 9 // só pode usar := dentro de code blocks

	fmt.Printf("A temperatura de ebulição da água é %.2f graus Fahrenheit.\n", tempF)
	fmt.Printf("A temperatura de ebulição da água é %.2f graus Celsius.\n", tempC)

}
