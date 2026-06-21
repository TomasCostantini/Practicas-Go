/*
Guia de ejercicios GO:
Ejercicio 17:
En una elección existen 3 candidatos. Por cada elector se tienen los siguientes datos: sexo (H o M) y el
partido por el cual votó (1 a 3). Ingresar “n” electores, determinar e imprimir:
• La cantidad de mujeres que votaron.
• La cantidad de varones que votaron al candidato 3.
• Candidato que ganó la elección, considerando que no hay empate.
*/
package main

import "fmt"

var (
	votantes, partido, p1, p2, p3, mujeres, varonestres int
	genero                                              string
)

func main() {
	fmt.Println("Ingrese la cantidad de electores: ")
	fmt.Scanln(&votantes)
	for i := 0; i < votantes; i++ {
		fmt.Println("Ingrese el partido elegido: ")
		fmt.Scanln(&partido)
		fmt.Println("Ingrese su genero (H O M): ")
		fmt.Scanln(&genero)

		if genero == "M" || genero == "m" {
			mujeres++
		}

		if partido == 1 {
			p1++

		} else if partido == 2 {
			p2++

		} else if partido == 3 {
			if genero == "M" || genero == "m" {
				varonestres++
			}
			p3++
		}
	}
	if p1 > p2 && p1 > p3 {
		fmt.Println("El 1er partido es el ganador")
	} else if p2 > p1 && p2 > p3 {
		fmt.Println("El 2do partido es el ganador")
	} else if p3 > p1 && p3 > p2 {
		fmt.Println("El 3er partido es el ganador")

	}
	fmt.Println("La cantidad de mujeres que votaron es: ", mujeres)
	fmt.Println("La cantidad de varones que votaron al 3er candidato es: ", varonestres)
	fmt.Println("Programa finalizado")
}
