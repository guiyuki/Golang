package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ping := make(chan bool, 1)
	pong := make(chan bool, 1)
	bola := make(chan bool, 1)

	go Bola(bola, ping, pong)
	go Ping(ping, pong, bola)
	go Pong(pong, ping, bola)

	bola <- true

	for {
		time.Sleep(2000)
	}
}

func Bola(b <-chan bool, cpi chan<- bool, cpo chan<- bool) {
	for {
		val := <-b
		n := Random()
		fmt.Println("Joga a bola!")
		if 1 == n {
			fmt.Printf("Pong joga! \n")
			time.Sleep(time.Second)
			cpo <- val
		} else {
			fmt.Printf("Ping joga! \n")
			time.Sleep(time.Second)
			cpi <- val
		}
	}
}

func Ping(chanR <-chan bool, chanE chan<- bool, bola chan<- bool) {
	ponto := 0
	for {
		val := <-chanR
		n := Random()
		if 40 > n {
			fmt.Printf("Ping perde! \n")
			time.Sleep(time.Second)
			bola <- val
			ponto--
			fmt.Printf("Pong: %v\n\n", ponto)
		} else {
			fmt.Printf("Ping devolve! \n")
			time.Sleep(200)
			chanE <- val
		}
	}
}

func Pong(chanR <-chan bool, chanE chan<- bool, bola chan<- bool) {
	ponto := 0
	for {
		n := Random()
		val := <-chanR
		if 60 > n {
			fmt.Printf("Pong perde! \n")
			time.Sleep(time.Second)
			bola <- val
			ponto--
			fmt.Printf("Pong: %v\n\n", ponto)
		} else {
			fmt.Printf("Pong devolve! \n")
			time.Sleep(200)
			chanE <- val
		}
	}
}

func Random() int {
	return rand.Intn(100)
}
