/*Guia de ejercicios Go.
Ejercicio 10:
Leer 8 números enteros y calcular el promedio. Imprimir el resultado. */

package main

import "fmt"

const (
	f = 8
)

func main() {

	var (
		suma, numero int
		promedio     float32
	)

	for i := 0; i < f; i++ {
		fmt.Println("Ingrese un numero:")
		fmt.Scanln(&numero)

		suma = suma + numero
	}
	promedio = float32(suma) / float32(f)
	fmt.Println("El promedio de todos los numeros es: ")
	fmt.Println(promedio)
}

