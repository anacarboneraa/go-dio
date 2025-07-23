package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria) {
	fmt.Printf("Área: %.2f\n", g.area())
	fmt.Printf("Perímetro: %.2f\n", g.perimetro())
}

func main() {
	q := quadrado{lado: 5}
	c := circulo{raio: 3}

	fmt.Println("Quadrado:")
	medir(q)

	fmt.Println("Círculo:")
	medir(c)
}
