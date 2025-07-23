package main

import "fmt"

var count = 100

func main() {
	for i := 0; i < count; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
		} else {
			continue
		}
	}
}
