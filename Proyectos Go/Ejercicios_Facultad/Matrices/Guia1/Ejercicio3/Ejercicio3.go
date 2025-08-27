/*
Guia de ejercicios GO.
Ejercicio 29:
Cargar una matriz A de 10x10 y crear una función que reciba como parámetro la matriz cargada y devuelva la
suma de los elementos pares y la suma de los impares. Imprimir los resultados.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	filas    = 10
	columnas = 10
)

var suma_pares, suma_impares int

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
	for i := 0; i < filas; i++ {
		for ii := 0; ii < columnas; ii++ {
			fmt.Printf("%4d", A[i][ii])
		}
		fmt.Println("")
	}
	ordenamiento(A)
	fmt.Println("La suma de los numeros pares es: ", suma_pares)
	fmt.Println("La suma de los numeros impares es: ", suma_impares)
}

func ordenamiento(A [filas][columnas]int) (sumapares, sumaimpares int) {
	for i := 0; i < filas; i++ {
		for ii := 0; ii < columnas; ii++ {
			if A[i][ii]%2 == 0 {
				suma_pares = suma_pares + A[i][ii]
			} else {
				suma_impares = suma_impares + A[i][ii]
			}
		}
	}
	return suma_pares, suma_impares
}
