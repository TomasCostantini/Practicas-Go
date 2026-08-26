/*
Guia de ejercicios GO:
Ejercicio 13:
Cargar una matriz A de 10x20 elementos. Generar una función que devuelva como resultado el mayor
elemento de las filas pares y el menor de las filas impares y las posiciones donde se ubican. La función deberá
recibir como parámetros:
a) La matriz a buscar
b) Si se desea buscar el mayor o menor (parámetro booleano)
c) Si se desea buscar sobre filas pares o impares (parámetro booleano)
Imprimir los resultados.
*/
package main

import "fmt"

const (
	filas    = 10
	columnas = 20
)

func main() {
	var (
		A      [filas][columnas]int
		numero int
		tipo   bool
	)

	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Println("Ingrese un numero: ")
			fmt.Scan(&numero)
			A[i][z] = numero
		}
	}
	fmt.Println("Ingrese el tipo de valor que desea buscar (0: mayor, 1:menor) ")
	fmt.Scan(&tipo)
	fmt.Println("Ingrese el tipo de valor que desea buscar (0: mayor, 1:menor) ")

}
