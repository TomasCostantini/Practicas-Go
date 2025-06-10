/*
Guia de ejercicios Go
Ejercicio 15:
Leer el nombre y la edad de 10 alumnos de un curso. Determinar e imprimir en pantalla:
• El nombre del alumno de mayor edad y el de menor edad.
• Cantidad de alumnos menores que 12.
• Cantidad de alumnos entre 10 y 16.
• El promedio de las edades.
*/
package main

import "fmt"

var (
	menores   int
	entre     int
	suma      int
	promedio  int
	edad      int
	edadmenor int
	edadmayor int

	nombremayor string
	nombremenor string
	nombre      string

	F = 10
)

func main() {

	for i := 0; i < F; i++ {

		fmt.Println("Ingrese el nombre del alumno:")
		fmt.Scanln(&nombre)
		fmt.Println("Ingrese la edad del alumno:")
		fmt.Scanln(&edad)

		if i == 0 {
			nombremayor = nombre
			nombremenor = nombre
			edadmenor = edad
			edadmayor = edad
		}

		if edad < edadmenor {
			edadmenor = edad
			nombremenor = nombre
		}

		if edad > edadmayor {
			edadmayor = edad
			nombremayor = nombre
		}

		if edad <= 16 && edad >= 13 {
			entre = entre + 1
		}
		if edad < 12 {
			menores = menores + 1
		}
		suma += edad
	}
	promedio = suma / F

	fmt.Println("El nombre y la edad del alumno mayor es: ", nombremayor, " ", edadmayor)
	fmt.Println("El nombre y la edad del alumno menor es: ", nombremenor, " ", edadmenor)
	fmt.Println("Cantidad de alumnos menores de 12: ", menores)
	fmt.Println("Cantidad de alumnos entre 10 y 16: ", entre)
	fmt.Println("Promedio de edades: ", promedio)

}
