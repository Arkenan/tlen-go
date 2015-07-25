package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Esto devuelve un canal al que se le manda el primer mensaje entre 2 canales.
func fanIn(input1, input2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// Boring crea un canal al cual se le manda permanentemente un mensaje.
func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%v: %v \n", i, msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 20; i++ {
		fmt.Printf("%v", <-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
