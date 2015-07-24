package main

import "fmt"

type sliceEnteros []int

func (s *sliceEnteros) agregar(elemento int) {
	l := len(*s)
	c := cap(*s)
	var sret sliceEnteros
	
	if l + 1 > c {
		sret = sliceEnteros(make([]int, l, 2*c))
		for i, v := range *s {
			sret[i] = v
		}
	} else {
		sret = *s
	}
	sret = sret[:l+1]
	sret[l] = elemento

	// Acá es donde se modifica el valor original.
	*s = sret
}

func main() {
	s := sliceEnteros([]int{1, 2, 3})
	fmt.Println(s, "tiene longitud", len(s), "y capacidad", cap(s))
	s.agregar(20)
	fmt.Println(s, "tiene longitud", len(s), "y capacidad", cap(s))
}
