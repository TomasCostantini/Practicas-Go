/*
Ejercicio de ordenamiento de vectores utilizando el metodo de la burbuja, a su vez, se utilizan funciones para
elegir si el ordenamiento es ascendente o descendente.
La variable "referencia" se utiliza para dar el valor de inicio del vector,
Mientras que la variable "referencia_maxima" se utiliza para dar el valor de fin del vector.
*/
package main

import "fmt"

var (
	vector            [10]int
	indice1           int
	indice2           int
	aux               int
	referencia        int = 0
	referencia_maxima int = 5
	opcion            int
)

func main() {
	for indice1 = referencia; indice1 < referencia_maxima; indice1++ {
		fmt.Print("Ingrese un numero: ")
		fmt.Scan(&vector[indice1])
	}

	fmt.Println("Seleccione el tipo de ordenamiento, 1 ascendente, 2 descendente: ")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		ordenamiento_ascendente()
	case 2:
		ordenamiento_descendente()
	default:
		fmt.Println("Opcion no valida")
	}

	fmt.Println("El vector ordenado es el siguiente:: ")
	for indice1 = referencia; indice1 < referencia_maxima; indice1++ {
		fmt.Println(vector[indice1])
	}
}

func ordenamiento_descendente() {
	for indice1 = referencia; indice1 < referencia_maxima; indice1++ {
		for indice2 = indice1 + 1; indice2 < referencia_maxima; indice2++ {
			if vector[indice1] < vector[indice2] {
				aux = vector[indice1]
				vector[indice1] = vector[indice2]
				vector[indice2] = aux
			}
		}
	}
}

func ordenamiento_ascendente() {
	for indice1 = referencia; indice1 < referencia_maxima; indice1++ {
		for indice2 = indice1 + 1; indice2 < referencia_maxima; indice2++ {
			if vector[indice1] > vector[indice2] {
				aux = vector[indice1]
				vector[indice1] = vector[indice2]
				vector[indice2] = aux
			}
		}
	}
}
