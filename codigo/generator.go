package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(c chan string) {
	//esto manda infinitamente mensajes a c.
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%v: hola que tal\n", i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func generar() chan string {
	c := make(chan string)
	go boring(c)
	return c
}

func main() {
	c := generar()
	//recibo los diez primeros mensajes ni bien estén disponibles.
	for i := 0; i < 10; i++ {
		println("El canal dice:", <-c)
	}
}
