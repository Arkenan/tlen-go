package main

import "fmt"

const c = 5
const c2 = 50000000000000000000

func funcionDeEnteros(n int) {
	fmt.Printf("La variable es de tipo %T y tiene valor: %v.\n", n, n)
}

func funcionDeFloats(f float32) {
	fmt.Printf("La variable es de tipo %T y tiene valor: %v.\n", f, f)
}

func main() {
	funcionDeEnteros(c)
	funcionDeFloats(c)
	funcionDeFloats(c2)
	//funcionDeEnteros(c2)
}
