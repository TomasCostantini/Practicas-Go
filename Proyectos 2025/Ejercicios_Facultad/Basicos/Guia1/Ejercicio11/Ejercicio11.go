/* Guia de ejercicios GO
Ejercicio 11:
Leer 10 números enteros y calcular la suma de los números pares y la suma de los impares. Imprimir el
resultado.
*/

package main

import "fmt"

func main() {

var( num int
	 sumaPares int
	 sumaImpares int
     F int = 10
	)

	
	for i := 0; i < F; i++ {
		fmt.Print("Ingrese un número entero: ")
		fmt.Scanln(&num)

		if num%2 == 0 {
			sumaPares += num
		} else {
			sumaImpares += num
		}
	}

	fmt.Println("La suma de los números pares es:", sumaPares)
	fmt.Println("La suma de los números impares es:", sumaImpares)
}