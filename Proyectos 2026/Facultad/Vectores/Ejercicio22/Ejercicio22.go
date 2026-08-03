/*Guia de ejercicios GO:
Ejercicio 22:
Cargar un vector de 5 números enteros e imprimir la suma y el promedio de dichos elementos.
*/

package main

import "fmt"

func main() {
	var (
		numeros  [5]int
		suma, n  int
		promedio float64
	)

	for i := 0; i < 5; i++ {
		fmt.Println("Ingrese un número: ")
		fmt.Scan(&n)
		numeros[i] = n
	}

	for i := 0; i < 5; i++ {
		suma += numeros[i]
	}
	promedio = float64(suma) / float64(5)
	fmt.Println("La suma es:", suma)
	fmt.Println("El promedio es:", promedio)
	fmt.Println("Programa finalizado")
}
