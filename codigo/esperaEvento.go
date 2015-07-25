package main

import "fmt"

func imprimir(n int, c chan int) {
	for i := 0; i < n; i++ {
		fmt.Println(n)
	}
	c <- 1 // aviso de fin.
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go imprimir(5-i, c)
	}

	for i := 0; i < 5; i++ {
		<-c // descarte.
	}
}
