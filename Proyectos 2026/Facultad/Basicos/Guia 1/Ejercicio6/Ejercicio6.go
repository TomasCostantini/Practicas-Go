/*
Guia de ejercicios GO:
Ejercicio 6:
Leer tres números enteros e imprimir el mayor de ellos.
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

	if numero1 > numero2 && numero1 > numero3 {
		fmt.Println("El numero mayor es: ", numero1)
	}
	if numero2 > numero1 && numero2 > numero3 {
		fmt.Println("El numero mayor es: ", numero2)
	}
	if numero3 > numero2 && numero3 > numero1 {
		fmt.Println("El numero mayor es: ", numero3)
	}
	fmt.Println("Programa finalizado.")
}
