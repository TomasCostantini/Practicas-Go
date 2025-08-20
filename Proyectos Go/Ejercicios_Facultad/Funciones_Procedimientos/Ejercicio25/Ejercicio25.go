/*
Guia de ejercicios GO.
Ejercicio 25:
Se tienen 2 vectores A y B de 20 elementos. Programar un procedimiento que permita ordenar
vectores de forma ascendente. El procedimiento debe recibir como parámetro por referencia el
vector a ordenar.
Una vez ordenado los 2 vectores mediante la llamada al procedimiento, generar un nuevo vector C
que sea la intercalación ordenada de A y de B (considere que no hay elementos repetidos).

los vectores A y B deben pasarse por referencia y cargarse en el main utilizando random
los vectores deben ordenarse con el metodo de la burbuja
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var vectorA [20]int
	var vectorB [20]int
	var vectorC [40]int

	cargarVector(&vectorA)
	cargarVector(&vectorB)
	ordenarVector(&vectorA)
	ordenarVector(&vectorB)
	intercalarVectores(&vectorA, &vectorB, &vectorC)

	fmt.Println("Vector A:", vectorA)
	fmt.Println("Vector B:", vectorB)
	fmt.Println("Vector C:", vectorC)
}

func cargarVector(vector *[20]int) {
	for i := 0; i < 20; i++ {
		vector[i] = rand.Intn(100)
	}
}

func ordenarVector(vector *[20]int) {
	for i := 0; i < len(*vector); i++ {
		for j := 0; j < len(*vector)-1; j++ {
			if (*vector)[j] > (*vector)[j+1] {
				(*vector)[j], (*vector)[j+1] = (*vector)[j+1], (*vector)[j]
			}
		}
	}
}
func intercalarVectores(vectorA, vectorB *[20]int, vectorC *[40]int) {
	i, j, k := 0, 0, 0
	for i < len(*vectorA) && j < len(*vectorB) {
		if (*vectorA)[i] < (*vectorB)[j] {
			(*vectorC)[k] = (*vectorA)[i]
			i++
		} else {
			(*vectorC)[k] = (*vectorB)[j]
			j++
		}
		k++
	}

	for i < len(*vectorA) {
		(*vectorC)[k] = (*vectorA)[i]
		i++
		k++
	}
	for j < len(*vectorB) {
		(*vectorC)[k] = (*vectorB)[j]
		j++
		k++
	}
}
