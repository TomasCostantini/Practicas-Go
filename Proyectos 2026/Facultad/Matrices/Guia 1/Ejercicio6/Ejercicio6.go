/*
Guia de ejercicios GO:
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
	var (
		numero int
		A      [filas][columnas]int
	)
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Println("Ingrese un numero: ")
			fmt.Scan(&numero)
			A[i][z] = numero
		}
	}
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			if i%2 == 0 {
				fmt.Print(A[i][z])
			}
		}
		fmt.Println("")
	}

}
