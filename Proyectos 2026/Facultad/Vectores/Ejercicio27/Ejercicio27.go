/*Guia de ejercicios GO:
Ejercicio 27:
Cargar un vector de 10 elementos, ordenarlo en forma ascendente e imprimirlo ordenado.
*/

package main

import "fmt"

func main() {
	const (
		indice = 10
	)
	var (
		vector [indice]int
		n, aux int
	)

	for i := 0; i < indice; i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scan(&n)
		vector[i] = n
	}

	for i := 0; i < indice; i++ {
		for z := 1; z < indice-1; z++ {
			if vector[i] < vector[z] {
				aux = vector[i]
				vector[i] = vector[z]
				vector[z] = aux
			}
		}
	}
	for i := 0; i < indice; i++ {
		fmt.Println(vector[i])
	}

	fmt.Println("Programa finalizado")
}
