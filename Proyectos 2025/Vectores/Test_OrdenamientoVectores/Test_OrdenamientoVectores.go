/*
Algoritmo de ordenamiento de vectores, metodo de la burbuja, se ingresan 10 numeros a un vector
y posteriormente se ordenan, el vector resultante queda ordenado de menor a mayor.
La variable "referencia" se utiliza para dar el valor de inicio del vector,
Mientras que la variable "referencia_maxima" se utiliza para dar el valor de fin del vector.
*/
package main

import "fmt"

func main() {
	var (
		vector            [10]int
		indice1           int
		indice2           int
		aux               int
		referencia        int = 0
		referencia_maxima int = 10
	)

	for indice1 = referencia; indice1 < referencia_maxima; indice1++ {
		fmt.Print("Ingrese un numero: ")
		fmt.Scan(&vector[indice1])
	}

	for indice1 = referencia; indice1 < referencia_maxima; indice1++ {
		for indice2 = indice1 + 1; indice2 < referencia_maxima; indice2++ {
			if vector[indice1] > vector[indice2] {
				aux = vector[indice1]
				vector[indice1] = vector[indice2]
				vector[indice2] = aux
			}
		}
	}

	fmt.Println("El vector ordenado es el siguiente:: ")
	for indice1 = referencia; indice1 < referencia_maxima; indice1++ {
		fmt.Println(vector[indice1])
	}
}
