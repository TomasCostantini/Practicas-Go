/*Guia de ejercicios GO:
Ejercicio 20:
Se lee la siguiente información: Sexo, Estado Civil y Edad en donde: sexo podrá tomar los valores H o M,
estado civil S (soltero), E (separado), C (casado) y V (viudo). Se pide:
• Contar la cantidad de hombres casados mayores de 28 años.
• Contar la cantidad de mujeres solteras menores de 25 años.
• Calcular el promedio de edad de los hombres casados.
• Contar cuantos hombres y mujeres se ingresaron.
• Contar la cantidad de hombres viudos
El final de la lectura viene indicado por un “S” en el campo sexo.
*/

package main

import "fmt"

var (
	sexo, estado_civil            string
	edad                          int
	casados_mayores               = 0
	mujeres_soteras_menores       = 0
	promedio_edad_hombres_casados = 0
	cantidad_hombres              = 0
	cantidad_mujeres              = 0
	hombres_viudos                = 0
	suma_edad_casados             = 0
)

func main() {
	fmt.Println("Ingrese el sexo de la persona (M o H): ")
	fmt.Scan(&sexo)
	for sexo != "S" || sexo != "s" {
		fmt.Println("Ingrese la edad de la persona: ")
		fmt.Scan(&edad)
		fmt.Println("Ingrese el estado civil de la persona(S:soltero, C:casado,V: viudo): ")
		fmt.Scan(&estado_civil)

		if sexo == "H" || sexo == "h" {
			cantidad_hombres++
			if estado_civil == "v" || estado_civil == "V" {
				hombres_viudos++
			}
			if estado_civil == "C" || estado_civil == "c" {
				suma_edad_casados = suma_edad_casados + edad
				if edad > 28 {
					casados_mayores++

				}

			} else {
				cantidad_mujeres++
				if edad > 25 && estado_civil == "S" || estado_civil == "s" {
					mujeres_soteras_menores++
				}
			}

		}
	}
	fmt.Println("La cantidad total de hombres es: ", cantidad_hombres)
	fmt.Println("La cantidad total de hombres casados mayores a 28 años es: ", casados_mayores)
	fmt.Println("La edad promedio de los hombres casados es: ", promedio_edad_hombres_casados)
	fmt.Println("La cantidad de hombre viudos es: ", hombres_viudos)
	fmt.Println("La cantidad total de mujeres es: ", cantidad_mujeres)
	fmt.Println("La cantidad total de mujeres solteras menores a 25 años es: ", mujeres_soteras_menores)
	fmt.Println("Programa finalizado")
}
