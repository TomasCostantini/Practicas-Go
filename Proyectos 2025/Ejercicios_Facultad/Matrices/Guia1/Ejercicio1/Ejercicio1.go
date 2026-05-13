/*
Guia de ejercicios Go.
Ejercicio 27:
Cargar una matriz A de 4x4. Crear una función que reciba como parámetro la matriz cargada y devuelva como
resultado la suma de sus elementos. Imprimir la suma.
*/
package main

import "fmt"

const (
	filas    = 4
	columnas = 4
)

func main() {
	var (
		A [filas][columnas]int
	)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print("Ingrese el elemento [", i, "][", j, "]: ")
			fmt.Scan(&A[i][j])
		}
	}
	fmt.Println("La suma de todos los elementos de la matriz es:", sumarElementos(A))
}

func sumarElementos(matriz [filas][columnas]int) int {
	var suma int
	for i := 0; i < filas; i++ {
		for j := 0; j < 4; j++ {
			suma += matriz[i][j]
		}
	}
	return suma
}
