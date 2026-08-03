/*Guia de ejercicios GO:
Ejercicio 24:
Cargar un vector de 7 números enteros y sumar aquellos elementos que sean distintos de su índice.
Imprimir la suma.
*/

package main

import "fmt"

func main() {
	var (
		numeros [7]int
		n       int
		suma    = 0
	)

	for i := 0; i < 6; i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scan(&n)
		numeros[i] = n
	}

	for i := 0; i < 6; i++ {
		if numeros[i] != i {
			suma = suma + numeros[i]
		}
	}
	fmt.Println("La suma total es: ", suma)
	fmt.Println("Programa finalizado")
}
