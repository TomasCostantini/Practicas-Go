/*Guia de ejercicios GO:
Ejercicio 28:
Cargar una matriz A de 6x5, Crear una función que reciba como parámetro la matriz cargada y devuelva la suma
de aquellos elementos que sean mayor o igual que 5 y la cantidad de elementos que intervinieron. Imprimir los
resultados.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	filas    = 6
	columnas = 5
)

func main() {
	var (
		A [filas][columnas]int
	)

	rand.Seed(time.Now().Unix())

	//Cargar la matriz
	for i := 0; i < filas; i++ {
		for ii := 0; ii < columnas; ii++ {
			A[i][ii] = rand.Intn(10)
		}
	}

	for f := 0; f < filas; f++ {
		for c := 0; c < columnas; c++ {
			fmt.Printf("%4d", A[f][c])
		}
		fmt.Println("")
	}
	fmt.Println("La suma de todos los valores es de: ", operacion(A))
}

func operacion(A [filas][columnas]int) (suma int) {
	for i := 0; i < filas; i++ {
		for ii := 0; ii < columnas; ii++ {
			if A[i][ii] >= 5 {
				suma = suma + A[i][ii]
			}
		}
	}
	return suma
}
