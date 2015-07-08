package main

import "fmt"

func countFrom(n int) func(int) int {
	sum := n
	return func(incr int) int {
		sum += incr
		return sum
	}
}

func main() {
	c1 := countFrom(10)
	c2 := countFrom(20)
	fmt.Println(c1(10), c2(10), c2(5))
}
