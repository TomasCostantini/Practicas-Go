/*
Guia de ejercicios GO
Ejercicio 6:
Cargar una matriz A de 4x4 elementos e imprimir aquellos elementos que se encuentren en filas pares.
*/
package main

import "fmt"

const (
	filas    = 4
	columnas = 4
)

func main() {
	var A [filas][columnas]int
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Print("Ingrese el elemento [", i, "][", j, "]: ")
			fmt.Scan(&A[i][j])
		}
	}
	fmt.Println("Elementos en filas pares:")
	for i := 0; i < filas; i++ {
		if i%2 == 0 {
			for j := 0; j < columnas; j++ {
				fmt.Print(A[i][j], " ")
			}
			fmt.Println()
		}
	}
}
