/*
Guia de ejercicios GO
Ejercicio 31:
Cargar una matriz de 6x6 elementos y poner un 0 en donde encuentre un valor par.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var matriz [6][6]int
	rand.Seed(time.Now().UnixNano())


	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			matriz[i][j] = rand.Intn(100) 
		}
	}
	fmt.Println("Matriz principal:")
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if matriz[i][j]%2 == 0 {
				matriz[i][j] = 0
			}
		}
	}

	fmt.Println("Matriz final:")
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}
