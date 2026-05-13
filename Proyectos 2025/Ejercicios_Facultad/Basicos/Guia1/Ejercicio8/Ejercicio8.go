/* Guia de ejercicios Go
Ejercicio 8
Leer tres datos: hora, minutos y segundos (string). Imprimir los mismos datos en el formato ‘hh:mm:ss’.
Ejemplo: “03:55:09”*/

package main

import "fmt"

func main() {
	var hora, minutos, segundos string
	fmt.Print("Ingrese la hora: ")
	fmt.Scan(&hora)
	fmt.Print("Ingrese los minutos: ")
	fmt.Scan(&minutos)
	fmt.Print("Ingrese los segundos: ")
	fmt.Scan(&segundos)

	fmt.Printf("La hora es: %s:%s:%s\n", hora, minutos, segundos)
}
