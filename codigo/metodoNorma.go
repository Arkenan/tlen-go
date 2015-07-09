package main

import (
	"fmt"
	"math"
)

type vertice struct {
	x int
	y int
	z int
}

func (v vertice) norma() float64 {
	return math.Sqrt(math.Pow(float64(v.x), 2) +
		math.Pow(float64(v.y), 2) +
		math.Pow(float64(v.z), 2))
}
func main() {
	p := vertice{3, 4, 5}
	fmt.Println(p.norma())
}
