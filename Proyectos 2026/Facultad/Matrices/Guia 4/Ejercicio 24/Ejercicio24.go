/* Guia de ejercicios GO:
Ejercicio 24:
Una empresa necesita confeccionar un programa que calcule el sueldo de sus empleados a partir de las horas
trabajadas de los mismos. Se sabe que los empleados son 67 y están identificados por el número de legajo en
forma correlativa, de 1 a 67. El sueldo se calcula de la siguiente manera: Sueldo = (Hs Normales * 10000 + Hs
Extras * 15000)

Cargar una matriz A con los siguientes datos: Nro. de legajo, Horas normales trabajadas, Horas Extras
trabajadas. Los legajos no necesariamente se leen en forma ordenada, pero en la matriz A deben guardarse en
la fila que corresponde, es decir, que el legajo 1 debe guardarse en la primer fila de la matriz, el legajo 2 en la
segundo, etc.

Se desea obtener por cada empleado el sueldo del mismo, guardar el dato en la misma matriz.
Ordenar la matriz por Sueldo de forma ascendente e imprimirla:
Nro. Legajo           Sueldo
0023                  70.000
0002                  55.000

Resolver el ejercicio en diagrama de flujo y luego programarlo en Go.
*/
package main

import "fmt"

const (
	empleados = 67
	filas     = 68
	columnas  = 4
)

func main() {
	var (
		sueldo, legajo, horas_normales, horas_extras int
		A                                            [filas][columnas]int
	)

	for i := 0; i < empleados; i++ {
		fmt.Println("Ingrese el numero de legajo del empleado: ")
		fmt.Scan(&legajo)
		A[legajo][1]=legajo
		fmt.Println("Ingrese las horas normales trabajadas: ")
		fmt.Scan(&horas_normales)
		A[legajo][3]=horas_normales
		fmt.Println("Ingrese la cantidad de horas extras trabajadas: ")
		fmt.Scan(&horas_extras)
		A[legajo][3]=horas_extras
	}
}
