package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NCPU = 4 // Cantidad de Núcleos.
type Vector []int
type VF []func(int) int

// Aplicar las operaciones de ops a v, de v[i] a v[n-1]
func (v Vector) DoSome(i, n int, ops VF, c chan int) {
	for ; i < n; i++ {
		v[i] += ops[i](v[i])
		fmt.Printf("v[%v] == %v (tarea %v) \n",
			i, v[i], i*NCPU/len(v)+1)
	}
	c <- 1 // Avisar que terminó esta tarea.
}

func (v Vector) DoAll(u VF) {
	c := make(chan int, NCPU) // Canal con un buffer por núcleo.
	for i := 0; i < NCPU; i++ {
		// Se separa en NCPU sectores al vector y se manda cada una a una goroutine.
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// Drain the channel.
	for i := 0; i < NCPU; i++ {
		<-c // Espera a que cada procesador termine.
	}
}

func crearTasks(n int) VF {
	vret := VF(make([]func(int) int, n))
	for i := 0; i < n; i++ {
		vret[i] = func(a int) int {
			cant := rand.Intn(1e3)
			t := time.Duration(cant) * time.Millisecond
			// Tardar un tiempo. Finge procesamiento.
			time.Sleep(t)
			// Devuelve la cantidad de milisegundos que tarda.
			return int(cant)
		}
	}
	return vret
}

func main() {
	vf := crearTasks(20)
	v1 := Vector(make([]int, 20))
	v1.DoAll(vf)
}
