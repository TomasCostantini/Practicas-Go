/*
Guia de ejercicios GO:
Ejercicio 5:
Cargar una matriz de 6x6 elementos y poner un 0 en donde encuentre un valor par.
*/
package main

import "fmt"

const (
	filas    = 6
	columnas = 6
)

func main() {
	var (
		numero int
		matriz [filas][columnas]int
	)
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Println("Ingrese un numero: ")
			fmt.Scan(&numero)
			matriz[i][z] = numero
		}
	}
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			if matriz[i][z]%2 == 0 {
				matriz[i][z] = 0
			}
		}
	}
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Print(matriz[i][z], "\t")
		}
		fmt.Println()
	}
	fmt.Println("Progrma finalizado")
}
