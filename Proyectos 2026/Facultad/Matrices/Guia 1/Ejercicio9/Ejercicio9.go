/*
Guia de ejercicios GO:
Ejercicio 9:
Cargar una matriz A de 5x3 e imprimir el menor elemento y la posición donde se encuentra (fila y columna).
*/
package main

import "fmt"

const (
	filas    = 5
	columnas = 3
)

func main() {
	var (
		menor, posicionfila, posicioncolumna, numero int
		A                                            [filas][columnas]int
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
				menor = A[i][z]
				posicionfila = i
				posicioncolumna = z
			} else {
					menor = A[i][z]
					posicionfila = i
					posicioncolumna = z
				}
			}
		}
	}

