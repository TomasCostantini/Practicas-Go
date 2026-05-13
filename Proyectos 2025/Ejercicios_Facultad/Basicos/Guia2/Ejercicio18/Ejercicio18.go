/*Teniendo como dato el número de matrícula y las notas de los 3 parciales evaluados, determinar si cada
alumno ingresado logró la condición de regularidad teniendo en cuenta:
• Las notas van del 1 al 10 y un parcial se aprueba con 7.
• Se regulariza la materia con al menos dos parciales aprobados.
• Un número de matrícula igual a 0 nos avisa el fin de datos.
Determinar además la cantidad total de alumnos, cuántos de ellos regularizaron la materia y el porcentaje
que éste grupo representa sobre el total. Imprimir los resultados.
*/

package main

import "fmt"

func main() {

	var (
		matricula     int
		nota1         float64
		nota2         float64
		nota3         float64
		alumnos       float64
		aprobados     int
		regularizados float64
		porcentaje    float64
	)

	fmt.Print("Ingrese el número de matrícula: ")
	fmt.Scan(&matricula)
	for matricula != 0 {
		fmt.Print("Ingrese la nota del primer parcial: ")
		fmt.Scan(&nota1)
		fmt.Print("Ingrese la nota del segundo parcial: ")
		fmt.Scan(&nota2)
		fmt.Print("Ingrese la nota del tercer parcial: ")
		fmt.Scan(&nota3)

		aprobados = 0

		if nota1 >= 7 {
			aprobados++
		}
		if nota2 >= 7 {
			aprobados++
		}
		if nota3 >= 7 {
			aprobados++
		}

		if aprobados >= 2 {
			regularizados++
		}
		alumnos++
		fmt.Print("Ingrese el número de matrícula: ")
		fmt.Scan(&matricula)
	}

	porcentaje = regularizados / alumnos * 100

	fmt.Println("Cantidad alumnos: ", alumnos)
	fmt.Println("Alumnos regularizados: ", regularizados)
	fmt.Println("Porcentaje de alumnos regularizados: ", porcentaje)

}
