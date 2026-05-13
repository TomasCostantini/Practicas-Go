/*Guia de ejercicios Go.
Ejercicio 1:
Calcular e imprimir la longitud y el área de una circunferencia a partir del radio leído (tipo real). Utilizar
constantes.
• longitud de la circunferencia = 2 * PI * radio
• área de la circunferencia = PI * radio * radio */

package main

import "fmt"

func main() {

	var (
		radio    float64
		area     float64
		longitud float64
	)
	const pi = 3.14

	fmt.Println("Ingresar el valor del radio: ")
	fmt.Scanln(&radio)

	longitud = 2 * pi * radio
	area = 2 * radio * radio

	fmt.Println("El valor de la longitud es: ", longitud)
	fmt.Println("El valor del area es: ", area)

}
