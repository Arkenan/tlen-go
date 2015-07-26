package main

import "fmt"

func main() {
	str := "Mirá!"
	fmt.Println("Runas:")
	for i, v := range str {
		fmt.Printf("%v: %q, que es de tipo %T\n", i, v, v)
	}
	fmt.Println("Bytes:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v: %q, que es de tipo %T\n", i, str[i], str[i])
	}
}
