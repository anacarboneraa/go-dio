package main

import (
	"fmt"
	"time"
)

func ping(pingChan chan bool, pongChan chan bool) {
	for {
		<- pingChan
		fmt.Println("ping")
		time.Sleep(1 * time.Second)
		pongChan <- true
	}
}

func pong(pingChan chan bool, pongChan chan bool) {
	for {
		<- pongChan
		fmt.Println("pong")
		time.Sleep(1 * time.Second)
		pingChan <- true
	}
}

func main() {
	pingChan := make(chan bool)
	pongChan := make(chan bool)

	go ping(pingChan, pongChan)
	go pong(pingChan, pongChan)

	pingChan <- true

	var enter string
	fmt.Scanln(&enter)
}
