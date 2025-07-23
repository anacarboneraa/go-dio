package main

import "fmt"

const count = 5

func main() {
	for i := 0; i < count; i++ {
		fmt.Println("Iteração:", i)
		if i % 2 == 0 {
			fmt.Println("Número é par")
		} else {
			fmt.Println("Número é ímpar")
		}
	}
}


