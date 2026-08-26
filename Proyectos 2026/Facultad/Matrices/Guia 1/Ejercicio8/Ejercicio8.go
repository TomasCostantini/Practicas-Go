/*
Guia de ejercicios GO:
Ejercicio 8:
Cargar una matriz A de 3x5 e imprimir el mayor elemento.
*/
package main

import "fmt"

const (
	filas    = 2
	columnas = 2
)

func main() {
	var (
		A             [filas][columnas]int
		mayor, numero int
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
			if i == 0 {
				mayor = A[i][z]
			} else {
				if A[i][z] > mayor {
					mayor = A[i][z]
				}
			}
		}
	}
	fmt.Println("El mayor elemento de la matriz es: ", mayor)
	fmt.Println("Progrma finalizado")
}
