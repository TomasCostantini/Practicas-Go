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
	votantes, partido, p1, p2, p3, mujeres, varonestres, genero int
)

func main() {
	fmt.Println("Ingrese la cantidad de electores: ")
	fmt.Scanln(&votantes)
	for i := 0; i < votantes; i++ {
		fmt.Println("Ingrese el partido elegido: ")
		fmt.Scanln(&partido)
		fmt.Println("Ingrese su genero: ")
		fmt.Scanln(&genero)
		if partido == 1 {
				if genero == 1 {

			}
			p1++
		} else if partido == 2 {
			if genero == 1 {

			}

		} else if partido == 3 {
			if genero == 1 {
				varonestres++
			}

		} else {
			fmt.Println("Numero no valido")
		}

	}
}
