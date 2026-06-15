/*Guia de ejercicios Go.
Ejercicio 13:
Leer 10 números enteros distintos y determinar cuál es el mayor y cual es el menor de ellos. Imprimir los
resultados.*/

package main

import "fmt"

const (
	f = 10
)

var (
	numero, mayor, menor int
)

func main() {

	for i := 0; i < f; i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scanln(&numero)

		if i == 1 {
			mayor = numero
			menor = numero
		} else {
			if numero > mayor {
				mayor = numero
			}
			if numero < menor {
				menor = numero
			}
		}

	}
	fmt.Println("El numero mayor es: ", mayor)
	fmt.Println("EL numero menor es: ", menor)
	fmt.Println("Programa finalizado")
}
