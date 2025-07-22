package main

import "fmt"

func main() {
	var nome string = "Ana Carbonera" // não precisa colocar o tipo da variavel, o Go infere automaticamente
	var idade int = 30
	var altura float32 = 1.65
	fmt.Println("Olá, meu nome é", nome, "tenho", idade, "anos e minha altura é", altura, "metros.")
	fmt.Printf("Olá, meu nome é %s, tenho %d anos e minha altura é %.2f metros.\n", nome, idade, altura)
}

// atrubuição :=

