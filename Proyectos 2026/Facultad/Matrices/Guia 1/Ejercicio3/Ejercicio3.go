/*
Guia de ejercicios GO:
Ejercicio 3:
Cargar una matriz A de 10x10 y crear una función que reciba como parámetro la matriz cargada y devuelva la
suma de los elementos pares y la suma de los impares. Imprimir los resultados.
*/
package main

import "fmt"

const (
	filas    = 2
	columnas = 2
)

func main() {
	var (
		A      [filas][columnas]int
		numero int
	)

	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Println("Ingrese un numero: ")
			fmt.Scan(&numero)
			A[i][z] = numero
		}
	}
	suma_pares, suma_impares := calculo(A)
	fmt.Println("La suma total de numeros impares es: ", suma_impares)
	fmt.Println("La suma total de numeros pares es: ", suma_pares)
}

func calculo(A [filas][columnas]int) (int, int) {
	var (
		suma_pares   = 0
		suma_impares = 0
	)

	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			if A[i][z]%2 == 0 {
				suma_pares = suma_pares + A[i][z]
			} else {
				suma_impares = suma_impares + A[i][z]
			}
		}
	}

	return suma_pares, suma_impares
}
