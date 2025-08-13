/*
Generador de numeros aleatorios, utilizando las librerias time y math/rand.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		numeroAleatorio := rand.Intn(100) // Genera un número aleatorio entre 0 y 99
		fmt.Println(numeroAleatorio)
	}
}
