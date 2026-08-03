/*Guia de ejercicios GO:
Ejercicio 28:
Se tienen las edades de 10 alumnos de un curso almacenadas en un vector ALU, se deberá realizar lo
siguiente:
a) Cargar el vector ALU
b) Encontrar la mayor y la menor edad del curso.
c) Encontrar la cantidad de estudiantes que se encuentran en las siguientes categorías:

• Menores o igual que 12
• entre 13 y 16 (inclusive)
• mayor o igual que 17
d) Ordenar el vector ALU de forma ascendente
*/

package main

import "fmt"

func main() {
	const (
		indice = 10
	)
	var (
		edad, mayor, menor, menordoce, entre, mayoritario, aux int
		ALU                                                    [indice]int
	)

	for i := 0; i < indice; i++ {
		fmt.Println("Ingrese la edad del alumno: ")
		fmt.Scan(&edad)
		ALU[i] = edad
	}

}
