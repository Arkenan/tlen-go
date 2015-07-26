package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%v: %v", i, msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("Hi! I'm Joe!")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(899 * time.Millisecond):
			fmt.Println("You're too slow.")
			return
		}
	}
}
