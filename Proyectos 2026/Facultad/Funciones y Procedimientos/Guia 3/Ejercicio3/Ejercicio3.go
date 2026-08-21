/*
Guia de ejercicios GO:
Ejercicio 3:
Se tienen 2 vectores A y B de 20 elementos. Programar un procedimiento que permita ordenar
vectores de forma ascendente (aplicando el método de la burbuja). El procedimiento debe recibir
como parámetro por referencia el vector a ordenar.
Una vez ordenado los 2 vectores mediante la llamada al procedimiento, generar un nuevo vector C
que sea la intercalación ordenada de A y de B (considere que no hay elementos repetidos).
*/

package main

import "fmt"

const (
	indice  = 4
	indiceC = 8
)

func main() {
	var (
		vectorA [indice]int
		vectorB [indice]int
		vectorC [indiceC]int
		numero  int
	)

	for i := 0; i < indice; i++ {
		fmt.Println("Ingrese el valor ", i, " del vector A:")
		fmt.Scan(&numero)
		vectorA[i] = numero
	}
	for i := 0; i < indice; i++ {
		fmt.Println("Ingrese el valor ", i, " del vector B:")
		fmt.Scan(&numero)
		vectorB[i] = numero
	}
	ordenamiento(&vectorA)
	ordenamiento(&vectorB)
	intercalar(&vectorA, &vectorB, &vectorC)
	fmt.Println(vectorC)
}

func ordenamiento(vector *[indice]int) {
	for i := 0; i < indice-1; i++ {
		for z := i + 1; i < indice; i++ {
			if vector[i] < vector[z] {
				aux := vector[i]
				vector[i] = vector[z]
				vector[z] = aux
			}
		}
	}
}

func intercalar(vectora *[indice]int, vectorb *[indice]int, vectorc *[indiceC]int) {
	for i := 0; i < indiceC; i = i + 2 {
		vectorc[i] = vectora[i]
	}
	for i := 1; i < indiceC; i = i + 2 {
		vectorc[i] = vectorb[i]
	}
}
