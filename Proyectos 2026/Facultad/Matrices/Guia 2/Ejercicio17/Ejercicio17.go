/*
Guia de ejercicios GO:
Ejercicio 17:
Cargar una matriz A de 5x5, ordenar cada fila en forma ascendente. Imprimir la matriz resultante.
*/
package main

import "fmt"

const (
	filas    = 5
	columnas = 5
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
}
