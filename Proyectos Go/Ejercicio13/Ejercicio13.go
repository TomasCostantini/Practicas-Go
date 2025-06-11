/*Guia de ejercicios GO:
Ejercicio 13:
Leer 10 números enteros distintos y determinar cuál es el mayor y cual es el menor de ellos. Imprimir los
resultados.
*/

package main

import "fmt"

func main() {

	var (
		num      int
		max      int
		min      int
		cantidad int = 10
	)

	for i := 0; i < cantidad; i++ {
		fmt.Print("Ingrese un número entero: ")
		fmt.Scanln(&num)
		if i == 0 {
			max = num
			min = num
		} else {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
	}

	fmt.Println("El número mayor es:", max)
	fmt.Println("El número menor es:", min)
}
