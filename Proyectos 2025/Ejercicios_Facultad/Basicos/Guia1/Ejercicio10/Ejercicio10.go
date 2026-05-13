/* Guia de ejercicios GO
Ejercicio 10:
Leer 8 números enteros y calcular el promedio. Imprimir el resultado.
*/

package main

import "fmt"

var (
	promedio int
	suma     int
	numero   int
	F        = 8
)

func main() {
	for i := 0; i < F; i++ {
		fmt.Println("Ingrese un numero entero: ")
		fmt.Scan(&numero)
		suma += numero
	}
	promedio = suma / F
	fmt.Println("El promedio de la suma de todos los valores es: ", promedio)
}
