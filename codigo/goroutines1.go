package main

import "fmt"

func imprimir(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(n)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go imprimir(i)
	}
}
