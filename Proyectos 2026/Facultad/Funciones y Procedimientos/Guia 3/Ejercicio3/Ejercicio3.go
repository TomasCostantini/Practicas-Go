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
		fmt.Println("Ingrese el valor ", i+1, " del vector A:")
		fmt.Scan(&numero)
		vectorA[i] = numero
	}
	for i := 0; i < indice; i++ {
		fmt.Println("Ingrese el valor ", i+1, " del vector B:")
		fmt.Scan(&numero)
		vectorB[i] = numero
	}
	ordenamiento(&vectorA)
	ordenamiento(&vectorB)
	intercalar(&vectorA, &vectorB, &vectorC)

	for i := 0; i < indice; i++ {
		fmt.Print(vectorA[i], ";")
	}
	fmt.Println()
	for i := 0; i < indice; i++ {
		fmt.Print(vectorB[i], ";")
	}
	fmt.Println()
	for i := 0; i < indiceC; i++ {
		fmt.Print(vectorC[i], ";")
	}
}

func ordenamiento(vector *[indice]int) {
	for i := 0; i < indice-1; i++ {
		for z := 0; z < indice-1-i; z++ {
			if vector[z] < vector[z+1] {
				aux := vector[z]
				vector[z] = vector[z+1]
				vector[z+1] = aux
			}
		}
	}
}

func intercalar(vectora *[indice]int, vectorb *[indice]int, vectorc *[indiceC]int) {

	for i := 0; i < indiceC; i += 2 {
		vectorc[i] = vectora[i/2]
	}

	for i := 1; i < indiceC; i += 2 {
		vectorc[i] = vectorb[i/2]
	}
}

// Intercala los elementos de vectorA y vectorB dentro de vectorC.
// El primer for coloca los valores de A en las posiciones pares de C,
// y el segundo coloca los valores de B en las posiciones impares.
// Se utiliza i/2 para obtener la posición correspondiente de A o B,
// ya que i avanza de dos en dos mientras los otros vectores avanzan de uno en uno.
