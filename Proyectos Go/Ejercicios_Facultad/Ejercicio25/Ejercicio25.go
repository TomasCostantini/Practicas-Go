/*Guia de ejercicios GO.
Ejercicio 25:
Se tienen 2 vectores A y B de 20 elementos. Programar un procedimiento que permita ordenar
vectores de forma ascendente. El procedimiento debe recibir como parámetro por referencia el
vector a ordenar.
Una vez ordenado los 2 vectores mediante la llamada al procedimiento, generar un nuevo vector C
que sea la intercalación ordenada de A y de B (considere que no hay elementos repetidos).


El procedimiento debe modificar la direccion de memoria de los vectores.
los vectores A y B deben pasarse por referencia y cargarse en el main utilizando random
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

	// Cargar vectores A y B con valores aleatorios
	cargarVector(&vectorA)
	cargarVector(&vectorB)

	// Ordenar vectores A y B
	ordenarVector(&vectorA)
	ordenarVector(&vectorB)

	// Generar vector C como intercalación de A y B
	generarVectorC(&vectorA, &vectorB, &vectorC)

	// Imprimir vector C
	fmt.Println("Vector C (intercalación de A y B):", vectorC)
}

func cargarVector(vector *[20]int) {
	for i := 0; i < 20; i++ {
		vector[i] = rand.Intn(100) // Generar números aleatorios entre 0 y 99
	}
}

func ordenarVector(vector *[20]int) {
	for i := 0; i < len(vector)-1; i++ {
		for j := i + 1; j < len(vector); j++ {
			if vector[i] > vector[j] {
				vector[i], vector[j] = vector[j], vector[i]
			}
		}
	}
}

func generarVectorC(vectorA, vectorB *[20]int, vectorC *[40]int) {
	i, j, k := 0, 0, 0
	for i < len(vectorA) && j < len(vectorB) {
		if vectorA[i] < vectorB[j] {
			vectorC[k] = vectorA[i]
			i++
		} else {
			vectorC[k] = vectorB[j]
			j++
		}
		k++
	}
	// Copiar elementos restantes de A o B
	for i < len(vectorA) {
		vectorC[k] = vectorA[i]
		i++
		k++
	}
	for j < len(vectorB) {
		vectorC[k] = vectorB[j]
		j++
		k++
	}
}
// Nota: Este código asume que no hay elementos repetidos en los vectores A y B,
// como se especifica en el enunciado del ejercicio.