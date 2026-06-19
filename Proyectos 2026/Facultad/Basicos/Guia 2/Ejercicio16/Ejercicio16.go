/*
Guia de ejercicios GO:
Ejercicio 16:
Leer los siguientes datos de los empleados de una empresa: Nombre, edad y sueldo. Determinar:
• el nombre de aquellos empleados que cumplan con la condición de que la edad > 25 y el sueldo
>3000.
• la cantidad de empleados menores de 18 años.
• la cantidad de empleados que se ingresaron y la suma de los sueldos.
El final de la lista viene dado por edad igual a cero. Imprimir todos los resultados en pantalla.
*/

package main

import "fmt"

var (
	nombre                                 string
	edad, sueldo, empleados, menores, suma int
)

func main() {
	fmt.Println("Ingrese la edad del empleado: ")
	fmt.Scanln(&edad)

	for edad > 0 {
		fmt.Println("Ingrese nombre del empleado: ")
		fmt.Scanln(&nombre)
		fmt.Println("Ingrese el sueldo del empleado: ")
		fmt.Scanln(&sueldo)

		if edad > 25 && sueldo > 3000 {
			fmt.Println("Cumple con las condiciones: ", nombre)
		}
		if edad < 18 {
			menores++
		}
		empleados++
		suma = suma + sueldo

		fmt.Println("Ingrese la edad del empleado: ")
		fmt.Scanln(&edad)
	}

	fmt.Println("La cantidad de empleados son: ", empleados)
	fmt.Println("La suma de los sueldos es de: ", suma, "pesos")
	fmt.Println("Programa finalizado")
}
