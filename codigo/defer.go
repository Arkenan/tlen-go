package main

import "fmt"

func main() {
	defer fmt.Printf("Mundo!\n")
	fmt.Print("Hola ")
}
