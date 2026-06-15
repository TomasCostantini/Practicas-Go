/*Guia de ejercicios GO:
Ejercicio 5:
Leer un número entero y determinar si es par o impar. (utilizar %)
*/

package main

import "fmt"

var numero int

func main() {
	fmt.Println("Ingrese un numero")
	fmt.Scanln(&numero)
	if numero%2 == 0 {
		fmt.Println("EL numero ingresado es par")
	} else {
		fmt.Println("El numero ingresado es impar")
	}
	fmt.Println("Programa finalizado")
}
