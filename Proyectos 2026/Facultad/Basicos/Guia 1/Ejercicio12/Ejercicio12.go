/*Guia de ejercicios Go.
Ejercicio 12:
Leer 10 números enteros y determinar cuántos son múltiplos de 5. Imprimir el resultado. (utilizar %) */

package main

import "fmt"

const (
	f = 10
)

func main() {

	var (
		numero int
	)

	for i := 0; i < f; i++ {
		fmt.Println("Ingrese un numero:")
		fmt.Scanln(&numero)
		if numero%5 == 0 {
			fmt.Println("Es multiplo de 5")
		} else {
			fmt.Println("No es multiplo de 5")
		}

	}
	fmt.Println("Programa finalizado")
}
