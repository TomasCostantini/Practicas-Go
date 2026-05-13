/*Guia de ejercicios Go
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

func main() {
	var (
		nombre            string
		edad              int
		sueldo            int
		empleados         string
		sueldototal       int
		contadorEmpleados int
		menores           int
	)

	fmt.Println("Ingrese el la edad del empleado: ")
	fmt.Scan(&edad)

	for edad != 0 {
		fmt.Println("Ingrese el nombre del empleado: ")
		fmt.Scan(&nombre)
		fmt.Println("Ingrese el sueldo del empleado: ")
		fmt.Scan(&sueldo)

		if edad > 25 && sueldo > 3000 {
			empleados += nombre + " , "
		}
		if edad < 18 {
			menores++
		}
		contadorEmpleados++
		sueldototal += sueldo

		fmt.Println("Ingrese la edad del empleado: ")
		fmt.Scan(&edad)
	}
	fmt.Println("Cantidad de empleados menores de 18 años: ", menores)
	fmt.Println("Cantidad de empleados ingresados: ", contadorEmpleados)
	fmt.Println("Suma de sueldos: ", sueldototal)
	fmt.Println("Empleados mayores de 25 años con sueldo mayor a 3000: ", empleados)
}
