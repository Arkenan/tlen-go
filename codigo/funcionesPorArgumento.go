package main

import "fmt"

func Map(arr []int, f func(int) int) []int {
	arr2 := make([]int, len(arr))
	for i := range arr {
		arr2[i] = f(arr[i])
	}
	return arr2
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	cuadrado := func(a int) int { return a * a }
	mitad := func(a int) int { return a / 2 }
	fmt.Println(Map(a, cuadrado))
	fmt.Println(Map(a, mitad))
}
