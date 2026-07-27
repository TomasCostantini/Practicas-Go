/*
Guia de ejercicios GO:
Ejercicio 18:
Teniendo como dato el número de matrícula y las notas de los 3 parciales evaluados, determinar si cada
alumno ingresado logró la condición de regularidad teniendo en cuenta:
• Las notas van del 1 al 10 y un parcial se aprueba con 7.
• Se regulariza la materia con al menos dos parciales aprobados.
• Un número de matrícula igual a 0 nos avisa el fin de datos.
Determinar además la cantidad total de alumnos, cuántos de ellos regularizaron la materia y el porcentaje
que éste grupo representa sobre el total. Imprimir los resultados.
*/
package main

import "fmt"

var (
	matricula, primernota, segundanota, tercernota, porcentaje int
	regulares                                                  = 0
	alumnos                                                    = 0
)

func main() {
	fmt.Println("Ingrese el numero de matricula: ")
	fmt.Scan(&matricula)
	for matricula > 0 {
		fmt.Println("Ingrese la primera nota: ")
		fmt.Scan(&primernota)
		fmt.Println("Ingrese la segunda nota: ")
		fmt.Scan(&segundanota)
		fmt.Println("Ingrese la tercer nota: ")
		fmt.Scan(&tercernota)

		if (primernota > 6 && segundanota > 6) || (primernota > 6 && tercernota > 6) || (segundanota > 6 && tercernota > 6) {
			fmt.Println("Alumno Regular")
			regulares++
		} else {
			fmt.Println("Alumno no aprobado")
		}
		alumnos++
		fmt.Println("Ingrese el numero de matricula: ")
		fmt.Scan(&matricula)
	}
	porcentaje = (regulares / alumnos) * 100
	fmt.Println("El total de alumnos es: ", alumnos)
	fmt.Println("La cantidad de alumnos regulares es ", regulares)
	fmt.Println("El prcentaje de alumnos aprobados es de: ", porcentaje)
	fmt.Println("Programa finalizado.")
}
