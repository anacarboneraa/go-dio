package main

import "fmt"

var count = 100

func main() {
	for i := 0; i < count; i++ {
		if i%3 == 0 {
			fmt.Print("Pin ")
		} else if i%5 == 0 {
			fmt.Print("Pan ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
