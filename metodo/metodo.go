package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

func (r retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2 * (r.comprimento + r.altura)
}

func main() {
	r := retangulo{comprimento: 10, altura: 5}

	fmt.Printf("Área do retângulo: %d\n", r.area())
	fmt.Printf("Perímetro do retângulo: %d\n", r.perimetro())
}
