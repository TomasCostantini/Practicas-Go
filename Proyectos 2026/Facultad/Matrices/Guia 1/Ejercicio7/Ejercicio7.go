/*
Guia de ejercicios GO:
Ejercicio 7:
Cargar una matriz A de 5x5 y contar aquellos elementos que coinciden con su fila o columna.
*/
package main

import "fmt"

const (
	filas    = 5
	columnas = 5
)

func main() {
	var (
		A        [filas][columnas]int
		contador = 0
		numero   int
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
			if A[i][z] == i || A[i][z] == z {
				contador++
			}
		}
	}
	fmt.Println("La cantidad de numeros iguales a las filas o columnas es: ", contador)

}
