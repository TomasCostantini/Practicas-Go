/*
Guia de ejercicios GO:
Ejercicio 1:
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
	fmt.Println("El valor de la suma de sus elementos es: ", sumar(A))
}

func sumar(A [filas][columnas]int) int {
	var suma int
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			suma = suma + A[i][z]
		}
	}
	return suma
}
