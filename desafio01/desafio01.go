package main

import "fmt"

const ebulicaoK = 373

func main() {

	var tempK = ebulicaoK
	tempC := tempK - 273

	fmt.Printf("A temperatura de ebulição da água é %d em Kelvin.\n", tempK)
	fmt.Printf("A temperatura de ebulição da água é %d graus Celsius.\n", tempC)

}
