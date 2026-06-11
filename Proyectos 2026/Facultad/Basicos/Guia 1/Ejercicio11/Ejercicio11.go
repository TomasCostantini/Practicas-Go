/*Guia de ejercicios Go.
Ejercicio 11:
Leer 10 números enteros y calcular la suma de los números pares y la suma de los impares. Imprimir el
resultado. */

package main

import "fmt"

const (
	f = 10
)

func main() {

	var (
		sumapar, sumaimpar, numero int
	)

	for i := 0; i < f; i++ {
		fmt.Println("Ingrese un numero:")
		fmt.Scanln(&numero)
		if numero%2 == 0 {
			sumapar = sumapar + numero

		} else {
			sumaimpar = sumaimpar + numero
		}

	}
	fmt.Println("El suma de todos los numeros pares es: ", sumapar)
	fmt.Println("La suma de todos los numeros impares es: ", sumaimpar)
}
