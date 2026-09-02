/*
Guia de ejercicios GO:
Ejercicio 16:
Cargar una matriz A de 4x4, intercambiar los elementos correspondientes a las filas pares con los de las filas
impares (la fila 1 con la 2; la 3 con la 4).
*/
package main

import "fmt"

const (
	filas    = 4
	columnas = 4
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
	fmt.Println("\nMatriz original:")
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Printf("%d\t", A[i][z])
		}
		fmt.Println()
	}

	for i := 0; i < filas; i += 2 {
		for z := 0; z < columnas; z++ {
			aux := A[i][z]
			A[i][z] = A[i+1][z]
			A[i+1][z] = aux
		}
	}

	fmt.Println("\nMatriz con las filas intercambiadas:")
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Printf("%d\t", A[i][z])
		}
		fmt.Println()
	}
}
