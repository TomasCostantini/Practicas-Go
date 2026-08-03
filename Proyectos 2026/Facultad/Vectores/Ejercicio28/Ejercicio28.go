/*Guia de ejercicios GO:
Ejercicio 28:
Se tienen las edades de 10 alumnos de un curso almacenadas en un vector ALU, se deberá realizar lo
siguiente:
a) Cargar el vector ALU
b) Encontrar la mayor y la menor edad del curso.
c) Encontrar la cantidad de estudiantes que se encuentran en las siguientes categorías:

• Menores o igual que 12
• entre 13 y 16 (inclusive)
• mayor o igual que 17
d) Ordenar el vector ALU de forma ascendente
*/

package main

import "fmt"

func main() {
	const (
		indice = 4
	)
	var (
		edad, mayor, menor, aux int
		ALU                     [indice]int
		menordoce               = 0
		mayoritario             = 0
		entre                   = 0
	)

	for i := 0; i < indice; i++ {
		fmt.Println("Ingrese la edad del alumno: ")
		fmt.Scan(&edad)
		ALU[i] = edad
	}

	for i := 0; i < indice; i++ {
		if i == 0 {
			mayor = ALU[i]
			menor = ALU[i]
		} else {
			if ALU[i] > mayor {
				mayor = ALU[i]
			}
			if ALU[i] < menor {
				mayor = ALU[i]
			}
		}
	}
	switch {
	case edad < 13:
		menordoce++
	case edad >= 13 && edad < 17:
		entre++
	case edad > 17:
		mayoritario++
	}
	for indiceA := 0; indiceA < indice; indiceA++ {
		for indiceB := indiceA + 1; indice < indice-1; indiceB++ {
			if ALU[indiceB] < ALU[indiceA] {
				aux = ALU[indiceB]
				ALU[indiceB] = ALU[indiceA]
				ALU[indiceA] = aux
			}
		}
	}

	fmt.Println("El alumno menor es: ", menor)
	fmt.Println("El alumno mayor es: ", mayor)
	fmt.Println("La cantidad de alumnos menores de 12 años es: ", menordoce)
	fmt.Println("La cantidad de alumnos con edad entre 13 y 16 es:", entre)
	fmt.Println("La cantidad de alumnos mayores a 17 años es: ", mayoritario)
	fmt.Println("El vector ordenado: ")
	for i := 0; i < indice; i++ {
		fmt.Print(ALU[i])
		fmt.Print(" , ")
	}
	fmt.Println("Programa finalizado")
}
