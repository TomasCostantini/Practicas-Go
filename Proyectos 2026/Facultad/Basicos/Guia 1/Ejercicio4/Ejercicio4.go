/*Guia de ejercicios GO:
Ejercicio 4:
Leer un número del 1 al 7 y determinar a qué día de la semana hace referencia. Imprimir el resultado
(lunes, martes, miércoles, jueves, viernes, sábado o domingo). Usar la estructura de selección múltiple
switch. */

package main

import "fmt"

var (
	dia int
)

func main() {
	fmt.Println("Ingrese el dia (lunes(1), martes(2), miércoles(3), jueves(4), viernes(5), sábado(6) o domingo(7))")
	fmt.Scanln(&dia)
	switch dia {
	case 1:
		fmt.Println("Es lunes")
	case 2:
		fmt.Println("Es martes")
	case 3:
		fmt.Println("Es miercoles")
	case 4:
		fmt.Println("Es jueves")
	case 5:
		fmt.Println("Es viernes")
	case 6:
		fmt.Println("Es sabado")
	case 7:
		fmt.Println("Es domingo")
	default:
		fmt.Println("El numero ingresado no es valido")
	}
	fmt.Println("Programa finalizado.")
}
