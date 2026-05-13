/* Guia de ejercicios Go
Ejercicio 3
Leer dos números enteros distintos. Imprimir en pantalla el mayor de ellos. */

package main

import "fmt"

func main() {
	var (
		primer_numero  int
		segundo_numero int
	)

	fmt.Println("Ingrese el valor del primer numero: ")
	fmt.Scanln(&primer_numero)
	fmt.Println("Ingrese el valor del segundo numero: ")
	fmt.Scanln(&segundo_numero)

	if primer_numero > segundo_numero {
		fmt.Println("El mayor numero es: ", primer_numero)
	} else {
		fmt.Println("El mayor numero es: ", segundo_numero)
	}

}
