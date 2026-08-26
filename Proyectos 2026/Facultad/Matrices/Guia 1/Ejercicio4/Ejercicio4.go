/*
Guia de ejercicios GO:
Ejercicio 4:
Cargar una matriz A de 5x5, luego solicitar al usuario un número N, contar la cantidad de veces que aparece N
en la matriz A, para ello crear una función que reciba como parámetro la matriz cargada y el número que se
desea contar. Imprimir el resultado.
*/
package main

import "fmt"

const (
	filas    = 5
	columnas = 5
)

func main() {
	var (
		numero, numero_referencia int
		A                         [filas][columnas]int
	)
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Println("Ingrese un numero: ")
			fmt.Scan(&numero)
			A[i][z] = numero
		}
	}
	fmt.Println("Ingrese un numero de referencia: ")
	fmt.Scan(&numero_referencia)
	fmt.Println("La cantidad de veces que se repite el numero es: ", conteo(&numero_referencia, &A))
	fmt.Println("Progrma finalizado")
}
func conteo(numero_referencia *int, A *[filas][columnas]int) int {
	var repeticiones = 0
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			if A[i][z] == *numero_referencia {
				repeticiones++
			}
		}
	}
	return repeticiones
}
