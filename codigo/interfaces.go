package main

import "fmt"

type persona string

type orco string

type camion string

func (persona) reir() {
	fmt.Println("JAJAJA")
}

func (orco) reir() {
	fmt.Println("kek")
}

func (c camion) andar() {
	fmt.Println("avanza", c)
}

type reidor interface {
	reir()
}

func main() {
	a := persona("Larry")
	b := orco("Thrall")
	c := camion("R-35")

	var i reidor
	i = a
	i.reir() // "JAJAJA"
	i = b
	i.reir() // "kek"
	i = c    // error en tiempo de compilación
}
