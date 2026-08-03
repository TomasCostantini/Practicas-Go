/*Guia de ejercicios GO:
Ejercicio 23:
Cargar un vector A de 10 elementos enteros. Ingresar un número entero N y contar la cantidad de veces
que aparece dicho número en el vector. Imprimir el resultado.
*/

package main

import "fmt"

func main() {
	var (
		vector    [10]int
		n, numero int
		cuenta    = 0
	)

	fmt.Println("Ingrese el numero de referencia: ")
	fmt.Scan(&n)
	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scan(&numero)
		vector[i] = numero
	}
	for i := 0; i < 10; i++ {
		if vector[i] == n {
			cuenta++
		}
	}
	fmt.Println("La cantidad de veces que n coincide es: ", cuenta)
	fmt.Println("Programa finalizado")
}
