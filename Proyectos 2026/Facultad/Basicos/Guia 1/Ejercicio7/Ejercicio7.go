/*
Guia de ejercicios GO:
Ejercicio 7:
Leer tres números enteros y determinar si se ingresaron en orden creciente o no. Escribir SI o NO según
corresponda.
*/
package main

import "fmt"

var numero1, numero2, numero3 int

func main() {
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scanln(&numero2)
	fmt.Println("Ingrese el tercer numero: ")
	fmt.Scanln(&numero3)

	if numero3 > numero2 && numero2 > numero1 {
		fmt.Println("SI")

	} else {
		fmt.Println("NO")
	}
	fmt.Println("Programa Finalizado")
}
