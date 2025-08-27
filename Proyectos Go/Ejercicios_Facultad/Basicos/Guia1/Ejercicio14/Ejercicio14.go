/*Guia de ejercicios Go
Ejercicio 14:
Leer un número entero e imprimir la suma de todos los números anteriores. Por ejemplo, si se ingresa el
número 4, imprimir el resultado de: 1 + 2 + 3.
*/
package main

import "fmt"

func main() {

	var (
		numero int
		suma   int
	)

	fmt.Print("Ingrese un número entero: ")
	fmt.Scanln(&numero)
	for i := 0; i < numero; i++ {
		suma += i
	}

	fmt.Println("La suma de todos los números anteriores a", numero, "es:", suma)
}
