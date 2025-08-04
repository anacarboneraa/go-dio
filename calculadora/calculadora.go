package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := subtrair(10, 5, 2)
	z := dividir(100, 2, 5)
	w := multiplicar(2, 3, 4)

	fmt.Println("Soma:", x)
	fmt.Println("Subtração:", y)
	fmt.Println("Divisão:", z)
	fmt.Println("Multiplicação:", w)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtrair(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

func dividir(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		if v == 0 {
			return 0
		}
		total /= v
	}
	return total
}

func multiplicar(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}
