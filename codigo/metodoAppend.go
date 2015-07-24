package main

import "fmt"

type sliceEnteros []int

func (s *sliceEnteros) agregar(elemento int) {
	var sret sliceEnteros
	if len(*s)+1 > cap(*s) {
		sret = sliceEnteros(make([]int, len(*s), 2*cap(*s)))
		for i, v := range *s {
			sret[i] = v
		}
	} else {
		sret = *s
	}
	sret = sret[:len(sret)+1]
	sret[len(sret)-1] = elemento

	// Acá es donde se modifica el valor original.
	*s = sret
}

func main() {
	s := sliceEnteros([]int{1, 2, 3})
	fmt.Println(s, "tiene longitud", len(s), "y capacidad", cap(s))
	s.agregar(20)
	fmt.Println(s, "tiene longitud", len(s), "y capacidad", cap(s))
}
