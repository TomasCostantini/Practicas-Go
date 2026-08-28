/*
Guia de ejercicios GO:
Ejercicio 10:
Cargar una matriz A de 4x4:
a) Crear una función que calcule y devuelva como resultado la suma de los elementos de la diagonal
principal
b) Crear una función que calcule y devuelva como resultado la suma de los elementos de la triangular
superior.
c) Imprimir los resultados.
*/
package main

import "fmt"

const (
	filas    = 4
	columnas = 4
)

func main() {
	var (
		A                                      [filas][columnas]int
		numero, suma_diagonal, suma_triangular int
	)

	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; i++ {
			fmt.Println("Ingrese un numero: ")
			fmt.Scan(&numero)
			A[i][z] = numero
		}
	}
	fmt.Println("La suma de los elemento de la diagonal principal es: ", suma_diagonal)
	fmt.Println("La suma de los elementos de la triangular superior es: ", suma_triangular)
}

func diagonal() {

}

func triangular() {

}
