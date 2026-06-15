/*
Guia de ejercicios GO:
Ejercicio 8:
Leer tres datos: hora, minutos y segundos (string). Imprimir los mismos datos en el formato ‘hh:mm:ss’.
Ejemplo: “03:55:09”
*/
package main

import "fmt"

var hora, minutos, segundos int

func main() {
fmt.Println("Ingrese las horas: ")
fmt.Scan(&hora)
fmt.Println("Ingrese los minutos: ")
fmt.Scan(&minutos)
fmt.Println("Ingrese los segundos: ")
fmt.Scan(&segundos)
fmt.Printf("%02d:%02d:%02d\n", hora, minutos, segundos)
}
