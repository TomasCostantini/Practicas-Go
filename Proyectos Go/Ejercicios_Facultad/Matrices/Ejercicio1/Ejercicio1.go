/*
Guia de ejercicios Go.
Ejercicio 1:
Cargar una matriz A de 4x4. Crear una función que reciba como parámetro la matriz cargada y devuelva como
resultado la suma de sus elementos. Imprimir la suma.
*/
package main

import "fmt"

func main() {
	var (
		A    [4][4]int
		suma int
	)
	
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print("Ingrese el elemento [", i, "][", j, "]: ")
			fmt.Scan(&A[i][j])
		}
	}
	suma = sumarElementos(A)
	fmt.Println("La suma de los elementos de la matriz es:", suma)
}

func sumarElementos(matriz [4][4]int) int {
	var suma int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			suma += matriz[i][j]
		}
	}
	return suma
}
