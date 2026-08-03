/*Guia de ejercicios GO:
Ejercicio 25:
Cargar un vector de 8 elementos enteros e imprimir el mayor de ellos.
*/

package main

import "fmt"

func main() {
	var (
		numeros  [8]int
		n, mayor int
	)

	for i := 0; i < 8; i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scan(&n)
		numeros[i] = n
	}

	for i := 0; i < 8; i++ {
		if i == 0 {
			mayor = numeros[i]
		} else {
			if numeros[i] > mayor {
				mayor = numeros[i]
			}
		}
	}
}
