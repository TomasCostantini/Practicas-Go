/*
Guia de ejercicios GO:
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
	nombre, nombremayor, nombremenor                               string
	edad, edadmayor, edadmenor, cantidadmenor, cantidadentre, suma int
	promedio                                                       float32
)

const (
	alumnos = 10
)

func main() {
	for i := 0; i < alumnos; i++ {
		fmt.Println("Ingrese el nombre del alumno: ")
		fmt.Scanln(&nombre)
		fmt.Println("Ingrese la edad del alumno: ")
		fmt.Scanln(&edad)

		if i == 1 {
			nombremayor = nombre
			edadmayor = edad
			nombremenor = nombre
			edadmenor = edad

		} else {
			if edad > edadmayor {
				nombremayor = nombre
				edadmayor = edad
			}
			if edad < edadmenor {
				nombremenor = nombre
				edadmenor = edad
			}
		}
		if edad < 12 {
			cantidadmenor++
		}
		if edad > 9 && edad < 17 {
			cantidadentre++
		}
		suma = suma + edad
	}
	promedio = float32(suma) / float32(alumnos)

	fmt.Println("El nombre del alumno de mayor edad es: ", nombremayor)
	fmt.Println("El nombre del alumno de menor edad es: ", nombremenor)
	fmt.Println("La cantidad de alumnos menores a 12 es: ", cantidadmenor)
	fmt.Println("La cantidad de alumnos entre 10 y 16 es: ", cantidadentre)
	fmt.Println("El promedio de edades es: ", promedio)
	fmt.Println("Programa finalizado")
}
