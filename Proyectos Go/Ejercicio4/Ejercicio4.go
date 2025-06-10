/* Guia de ejercicios Go
Ejercicio 4
Leer un número del 1 al 7 y determinar a qué día de la semana hace referencia. Imprimir el resultado
(lunes, martes, miércoles, jueves, viernes, sábado o domingo). Usar la estructura de selección múltiple
switch. */

package main

import "fmt"

var numero int

func main() {
	fmt.Println("Leer un numero del 1 al 7: ")
	fmt.Scan(&numero)

	switch numero {
	case 1:
		fmt.Println("Es Lunes")
	case 2:
		fmt.Println("Es Martes")
	case 3:
		fmt.Println("Es Miercoles")
	case 4:
		fmt.Println("Es jueves")
	case 5:
		fmt.Println("Es viernes")
	case 6:
		fmt.Println("Es Sabado")
	case 7:
		fmt.Println("Es Domingo")
	default:
		fmt.Println("El número no es valido, ingresar valores del 1 al 7")
	}

}
