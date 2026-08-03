/*Guia de ejercicios GO:
Ejercicio 26:
Cargar un vector A de 10 números enteros y generar un vector B que contenga solamente los números
positivos de A. Imprimir el vector generado.
*/

package main

import "fmt"

func main() {
	var (
		vectorA [10]int
		vectorB []int
		n       int
		indiceB = 0
	)
	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scan(&n)
		vectorA[i] = n
	}
	for i := 0; i < 10; i++ {
		if vectorA[i] > 0 {
			vectorB[indiceB] = vectorA[i]
			indiceB++
		}
	}

	for i := 0; i < indiceB; i++ {
		fmt.Print(vectorB[i])
	}
}
