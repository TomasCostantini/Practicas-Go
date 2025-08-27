/*
Guia de ejercicios Go
Ejercicio 30:
Cargar una matriz A de 5x5 con valores random, luego solicitar al usuario un número N, contar la cantidad de veces que aparece N
en la matriz A, para ello crear una función que reciba como parámetro la matriz cargada y el número que se
desea contar. Imprimir el resultado.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var matriz [5][5]int
	var n, contador int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matriz[i][j] = rand.Intn(20)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Print("Ingrese el numero a contar: ")
	fmt.Scan(&n)
	contador = contarApariciones(matriz, n)
	fmt.Println("El numero", n, "aparece", contador, "veces en la matriz.")
}

func contarApariciones(matriz [5][5]int, n int) int {
	contador := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matriz[i][j] == n {
				contador++
			}
		}
	}
	return contador
}
