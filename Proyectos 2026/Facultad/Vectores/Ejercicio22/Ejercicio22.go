/*Guia de ejercicios GO:
Ejercicio 22:
Cargar un vector de 5 números enteros e imprimir la suma y el promedio de dichos elementos.
*/

package main

import "fmt"

func main() {
	var numeros [5]int
	var suma int
	var promedio float64

	for i := 0; i < len(numeros); i++ {
		fmt.Printf("Ingrese el número %d: ", i+1)
		fmt.Scan(&numeros[i])
	}
	for i := 0; i < len(numeros); i++ {
		suma += numeros[i]
	}
	promedio = float64(suma) / float64(len(numeros))
	fmt.Println("La suma es:", suma)
	fmt.Printf("El promedio es: %.2f\n", promedio)
}
