/*Guia de ejercicios GO:
Ejercicio 3:
Leer dos números enteros distintos. Imprimir en pantalla el mayor de ellos. */

package main

import "fmt"

var (
	primernumero, segundonumero int
)

func main() {
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scanln(&primernumero)
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scanln(&segundonumero)

	if primernumero > segundonumero {
		fmt.Println("El mayor numero es: ", primernumero)
	} else {
		fmt.Println("El mayor numero es :", segundonumero)
	}

	fmt.Println("Programa finalizado.")

}
